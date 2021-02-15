package resolvers

import (
	"resident-graphql/connection"
	"resident-graphql/models/responses"
	"resident-graphql/models/tables"

	"github.com/graphql-go/graphql"
)

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
