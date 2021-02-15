package resolvers

import (
	"fmt"
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

// UpdateOfficeResolver func
func UpdateOfficeResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	id, _ := p.Args["id"].(int)
	name, _ := p.Args["office_name"].(string)

	var office tables.Offices

	db.Model(&office).Where("id = ?", id).Update("office_name", name)

	return office, nil
}

// DeleteOfficeResolver func
func DeleteOfficeResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	id, _ := p.Args["id"].(int)

	var office tables.Offices

	db.Delete(&office, id)
	fmt.Println(office)

	return office, nil
}
