package migration

import (
	"log"
	"resident-graphql/connection"
	"resident-graphql/models/tables"
)

// Migrations func to migrate db
func Migrations() {
	db := *connection.GetConnection()
	var (
		checkProvinces             bool
		checkDistricts             bool
		checkSubDistricts          bool
		checkPersons               bool
		checkUsers                 bool
		checkOffices               bool
		checkOfficePersonLocations bool
		checkQueue                 bool
	)

	db.Migrator().DropTable(&tables.Provinces{})
	db.Migrator().DropTable(&tables.Districts{})
	db.Migrator().DropTable(&tables.SubDistricts{})
	db.Migrator().DropTable(&tables.Persons{})
	db.Migrator().DropTable(&tables.Users{})
	db.Migrator().DropTable(&tables.Offices{})
	db.Migrator().DropTable(&tables.OfficePersonLocations{})

	checkProvinces = db.Migrator().HasTable(&tables.Provinces{})
	if !checkProvinces {
		db.Migrator().CreateTable(&tables.Provinces{})
		log.Println("Create Provinces Table")
	}

	checkDistricts = db.Migrator().HasTable(&tables.Districts{})
	if !checkDistricts {
		db.Migrator().CreateTable(&tables.Districts{})
		log.Println("Create Districts Table")
	}

	checkSubDistricts = db.Migrator().HasTable(&tables.SubDistricts{})
	if !checkSubDistricts {
		db.Migrator().CreateTable(&tables.SubDistricts{})
		log.Println("Create SubDistricts Table")
	}

	checkPersons = db.Migrator().HasTable(&tables.Persons{})
	if !checkPersons {
		db.Migrator().CreateTable(&tables.Persons{})
		log.Println("Create Persons Table")
	}

	checkUsers = db.Migrator().HasTable(&tables.Users{})
	if !checkUsers {
		db.Migrator().CreateTable(&tables.Users{})
		log.Println("Create Users Table")
	}

	checkOffices = db.Migrator().HasTable(&tables.Offices{})
	if !checkOffices {
		db.Migrator().CreateTable(&tables.Offices{})
		log.Println("Create Offices Table")
	}

	checkOfficePersonLocations = db.Migrator().HasTable(&tables.OfficePersonLocations{})
	if !checkOfficePersonLocations {
		db.Migrator().CreateTable(&tables.OfficePersonLocations{})
		log.Println("Create OfficePersonLocations Table")
	}

	checkQueue = db.Migrator().HasTable(&tables.QueueEmail{})
	if !checkQueue {
		db.Migrator().CreateTable(&tables.QueueEmail{})
		log.Println("Create QueueEmail Table")
	}
}
