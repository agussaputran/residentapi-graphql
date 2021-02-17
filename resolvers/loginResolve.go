package resolvers

import (
	"log"
	"os"
	"resident-graphql/connection"
	"resident-graphql/models/tables"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

// LoginResolver func
func LoginResolver(p graphql.ResolveParams) (interface{}, error) {
	var (
		db     = *connection.GetConnection()
		user   tables.Users
		result interface{}
	)

	email, _ := p.Args["email"].(string)
	password, _ := p.Args["password"].(string)

	db.Where("email = ?", email).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println("Email ", user.Email, " Password salah")
		result = map[string]interface{}{
			"message": "email atau password salah",
		}
	} else {
		type authCustomClaims struct {
			Email string `json:"email"`
			Role  string `json:"role"`
			jwt.StandardClaims
		}

		claims := &authCustomClaims{
			user.Email,
			user.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}
		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		token, err := sign.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			log.Println("Gagal create token, message ", err.Error())
			result = map[string]interface{}{
				"token": nil,
			}
		} else {
			log.Println("Email ", user.Email, " Berhasil login")
			result = map[string]interface{}{
				"email": user.Email,
				"token": token,
			}
		}
	}
	return result, nil
}
