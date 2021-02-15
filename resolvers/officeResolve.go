package resolvers

import (
	"math/rand"
	"resident-graphql/connection"
	"resident-graphql/models/responses"
	"resident-graphql/models/tables"
	"time"

	"github.com/graphql-go/graphql"
)

// CreateOfficeResolver func
func CreateOfficeResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	rand.Seed(time.Now().UnixNano())

	var office tables.Offices

	office.ID = uint(rand.Intn(100000))
	office.OfficeName = p.Args["office_name"].(string)
	office.SubDistrictID = uint(p.Args["sub_district_id"].(int))

	db.Create(&office)

	return office, nil
}

// ReadOfficeResolver func
func ReadOfficeResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	var (
		office   []tables.Offices
		response []responses.OfficeResponse
	)
	db.Model(&office).
		Select("offices.id, offices.office_name, sub_districts.sub_district_name").
		Joins("left join sub_districts on sub_districts.id = offices.sub_district_id").
		Scan(&response)
	return response, nil
}
