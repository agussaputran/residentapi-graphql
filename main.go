package main

import (
	"resident-graphql/connection"
	"resident-graphql/database/migration"
	"resident-graphql/database/seeders"
)

func main() {
	connection.Connect()
	migration.Migrations()
	seeders.Seeder()
}
