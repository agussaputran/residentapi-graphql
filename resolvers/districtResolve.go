package resolvers

import (
	"resident-graphql/connection"
	"resident-graphql/models/responses"
	"resident-graphql/models/tables"

	"github.com/graphql-go/graphql"
)

// ReadDistrictResolver func
func ReadDistrictResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	var (
		district []tables.Districts
		response []responses.DistrictResponse
	)
	db.Model(&district).
		Select("districts.id, districts.name as district, provinces.name as province").
		Joins("left join provinces on provinces.id = districts.province_id").Scan(&response)
	return response, nil
}
