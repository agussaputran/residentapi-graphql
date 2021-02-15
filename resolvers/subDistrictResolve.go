package resolvers

import (
	"resident-graphql/connection"
	"resident-graphql/models/responses"
	"resident-graphql/models/tables"

	"github.com/graphql-go/graphql"
)

// ReadSubDistrictResolver func
func ReadSubDistrictResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	var (
		subDistrict []tables.SubDistricts
		response    []responses.SubDistrictResponse
	)
	db.Model(&subDistrict).
		Select(`sub_districts.id, sub_districts.sub_district_name,
		districts.district_name, provinces.province_name`).
		Joins("left join provinces on provinces.id = districts.province_id").Scan(&response)
	return response, nil
}
