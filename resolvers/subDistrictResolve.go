package resolvers

import (
	"math/rand"
	"resident-graphql/connection"
	"resident-graphql/models/responses"
	"resident-graphql/models/tables"
	"time"

	"github.com/graphql-go/graphql"
)

// CreateSubDistrictResolver func
func CreateSubDistrictResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	rand.Seed(time.Now().UnixNano())

	var subDistrict tables.SubDistricts

	subDistrict.ID = uint(rand.Intn(100000))
	subDistrict.SubDistrictName = p.Args["sub_district_name"].(string)
	subDistrict.DistrictID = uint(p.Args["district_id"].(int))

	db.Create(&subDistrict)

	return subDistrict, nil
}

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
		Joins(`left join districts
		on districts.id = sub_districts.district_id left join provinces
		on provinces.id = districts.province_id`).Scan(&response)
	return response, nil
}

// UpdateSubDistrictResolver func
func UpdateSubDistrictResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	id, _ := p.Args["id"].(int)
	name, _ := p.Args["sub_district_name"].(string)

	var subDistrict tables.SubDistricts

	db.Model(&subDistrict).Where("id = ?", id).Update("sub_district_name", name)

	return subDistrict, nil
}
