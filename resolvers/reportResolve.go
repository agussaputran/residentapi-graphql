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
		Select().
		Joins().
		Scan(&response)
	return response, nil
}
