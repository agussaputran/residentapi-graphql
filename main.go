package main

import (
	"log"
	"net/http"
	"resident-graphql/api"
	"resident-graphql/connection"
	"resident-graphql/helper"
	"resident-graphql/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	connection.Connect()
	// migration.Migrations()
	// seeders.Seeder()

	app := gin.Default()
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
