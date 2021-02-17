package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"resident-graphql/api"
	"resident-graphql/connection"
	"resident-graphql/database/migration"
	"resident-graphql/database/seeders"
	"resident-graphql/helper"
	"resident-graphql/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	connection.Connect(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL"), os.Getenv("DB_TIMEZONE"))
	migration.Migrations()
	seeders.Seeder()

	app := gin.Default()
	app.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello docker",
		})
	})
	app.POST("/", func(c *gin.Context) {
		type Query struct {
			Query string
		}

		helper.Token = c.Request.Header.Get("Authorization")
		helper.ReqBody = middlewares.LogRequestBody(c)

		log.Println(middlewares.LogRequest(c))

		var query Query
		c.Bind(&query)

		result := api.ExecuteQuery(query.Query, api.Schema)
		c.JSON(http.StatusOK, result)
	})
	app.Run(":8080")
}
