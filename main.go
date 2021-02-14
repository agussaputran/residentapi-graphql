package main

import (
	"net/http"
	"resident-graphql/api"
	"resident-graphql/connection"

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

		var query Query
		c.Bind(&query)

		result := api.ExecuteQuery(query.Query, api.Schema)
		c.JSON(http.StatusOK, result)
	})
	app.Run(":8080")
}
