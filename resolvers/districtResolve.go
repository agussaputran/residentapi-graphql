package resolvers

import (
	"math/rand"
	"resident-graphql/connection"
	"resident-graphql/models/responses"
	"resident-graphql/models/tables"
	"time"

	"github.com/graphql-go/graphql"
)

// CreateDistrictResolver func
func CreateDistrictResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	rand.Seed(time.Now().UnixNano())

	var district tables.Districts

	district.ID = uint(rand.Intn(100000))
	district.DistrictName = p.Args["district_name"].(string)
	district.ProvinceID = uint(p.Args["province_id"].(int))

	db.Create(&district)

	return district, nil
}

// ReadDistrictResolver func
func ReadDistrictResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	var (
		district []tables.Districts
		response []responses.DistrictResponse
	)
	db.Model(&district).
		Select("districts.id, districts.district_name, provinces.province_name").
		Joins("left join provinces on provinces.id = districts.province_id").Scan(&response)
	return response, nil
}
