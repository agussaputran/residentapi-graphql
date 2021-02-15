package resolvers

import (
	"resident-graphql/connection"
	"resident-graphql/models/responses"
	"resident-graphql/models/tables"

	"github.com/graphql-go/graphql"
)

// ReadReportGenderResolver func
func ReadReportGenderResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	var (
		person   []tables.Persons
		response []responses.ReportGenderResponse
	)
	db.Model(&person).
		Select("gender, count(*) as total").Group("gender").Scan(&response)
	return response, nil
}

// ReadReportPersonOfficeByGenderResolver func
func ReadReportPersonOfficeByGenderResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	var (
		person   []tables.Persons
		response []responses.ReportGenderResponse
	)
	db.Model(&person).
		Select("gender, count(*) as total").
		Joins("join office_person_locations on persons.id = office_person_locations.person_id").
		Group("gender").Scan(&response)
	return response, nil
}
