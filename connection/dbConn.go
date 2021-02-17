package connection

import (
	"log"
	"os"
	"resident-graphql/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// Connect to DB
func Connect() {
	helper.SetEnvVar()
	var (
		dbHost     string
		dbUser     string
		dbPassword string
		dbName     string
		dbPort     string
		dbSsl      string
		dbTimeZone string
	)

	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbPort = os.Getenv("DB_PORT")
	dbSsl = os.Getenv("DB_SSL")
	dbTimeZone = os.Getenv("DB_TIMEZONE")

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
