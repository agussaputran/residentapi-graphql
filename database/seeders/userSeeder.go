package seeders

import (
	"log"
	"resident-graphql/models/tables"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seedUser(db *gorm.DB) {
	var userArray = [...][3]string{
		{"admin@residentapi.com", "admin", "admin"},
		{"data@residentapi.com", "data", "entry"},
		{"agussaputra@gmail.com", "agussaputra", "guest"},
	}

	var user tables.Users
	for _, v1 := range userArray {
		user.Email = v1[0]
		user.Password = v1[1]
		user.Role = v1[2]
		user.ID = 0

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error -> ", err.Error())
		}
		user.Password = string(hash)

		db.Create(&user)

	}
	log.Println("Seeder User created")
}
