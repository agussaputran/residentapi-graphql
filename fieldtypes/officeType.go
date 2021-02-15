package fieldtypes

import "github.com/graphql-go/graphql"

// OfficeType type
func OfficeType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Office",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"office_name": &graphql.Field{
					Type: graphql.String,
				},
				"sub_district_id": &graphql.Field{
					Type: graphql.Int,
				},
				"sub_district_name": &graphql.Field{
					Type: graphql.String,
				},
				"district_name": &graphql.Field{
					Type: graphql.String,
				},
				"province_name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
