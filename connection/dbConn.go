package connection

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// Connect to DB
func Connect(dbHost, dbUser, dbPassword, dbName, dbPort, dbSsl, dbTimeZone string) {

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
