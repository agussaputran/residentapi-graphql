package resolvers

import (
	"resident-graphql/connection"
	"resident-graphql/models/responses"
	"resident-graphql/models/tables"

	"github.com/graphql-go/graphql"
)

// ReadPersonResolver func
func ReadPersonResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	var (
		person   []tables.Persons
		response []responses.PersonResponse
	)
	db.Model(&person).
		Select(`persons.id, persons.nip, persons.full_name, persons.first_name,
		persons.last_name, persons.birth_date, persons.birth_place, persons.photo_profile_url,
		persons.gender, persons.zone_location, sub_districts.sub_district_name`).
		Joins("left join sub_districts on sub_districts.id = persons.sub_district_id").Scan(&response)
	return response, nil
}
