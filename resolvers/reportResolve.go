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

// ReadReportPersonOffice func
func ReadReportPersonOffice(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()

	var (
		response []responses.ReportPersonOfficeResponse
	)

	db.Table("persons").Select("persons.id as id, persons.full_name, count(*) as total").
		Joins(`join office_person_locations opl on opl.person_id = persons.id
				join offices ofc on opl.office_id = ofc.id
				join sub_districts sd on sd.id = ofc.sub_district_id
				join districts d on d.id = sd.district_id`).Group("persons.id, persons.full_name").Scan(&response)

	return response, nil
}
