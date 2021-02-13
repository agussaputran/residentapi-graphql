package seeders

import "resident-graphql/connection"

// Seeder func to seed table
func Seeder() {
	db := *connection.GetConnection()

	seedProvince(&db)
	seedDistrict(&db)
	seedSubDistrict(&db)
	seedPerson(&db)
	seedOffice(&db)
	seedOfficePersonLocation(&db)
	seedUser(&db)
}
