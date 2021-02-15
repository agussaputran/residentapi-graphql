package fieldtypes

import "github.com/graphql-go/graphql"

// ReportGenderType type
func ReportGenderType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Report",
			Fields: graphql.Fields{
				"gender": &graphql.Field{
					Type: graphql.String,
				},
				"total": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}

// ReportPersonOfficeType type
func ReportPersonOfficeType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Report",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"full_name": &graphql.Field{
					Type: graphql.String,
				},
				"total": &graphql.Field{
					Type: graphql.Int,
				},
				"districts": &graphql.Field{
					Type: graphql.NewList(DistrictType()),
				},
			},
		},
	)
}
