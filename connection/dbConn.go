package connection

import (
	"log"
	"resident-graphql/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// Connect to DB
func Connect() {
	var (
		dbHost     string
		dbUser     string
		dbPassword string
		dbName     string
		dbPort     string
		dbSsl      string
		dbTimeZone string
	)

	dbHost = helper.GetEnvVar("DB_HOST")
	dbUser = helper.GetEnvVar("DB_USER")
	dbPassword = helper.GetEnvVar("DB_PASSWORD")
	dbName = helper.GetEnvVar("DB_NAME")
	dbPort = helper.GetEnvVar("DB_PORT")
	dbSsl = helper.GetEnvVar("DB_SSL")
	dbTimeZone = helper.GetEnvVar("DB_TIMEZONE")

	conn := " host=" + dbHost +
		" user=" + dbUser +
		" password=" + dbPassword +
		" dbname=" + dbName +
		" port=" + dbPort +
		" sslmode=" + dbSsl +
		" TimeZone=" + dbTimeZone

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to postgreDB")
	} else {
		log.Println("successful connection")
	}

	dbConn = db
}

//GetConnection from DB
func GetConnection() *gorm.DB {
	return dbConn
}
