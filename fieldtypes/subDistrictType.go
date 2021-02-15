package fieldtypes

import "github.com/graphql-go/graphql"

// SubDistrictName type
func SubDistrictName() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "District",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},

				"sub_district_name": &graphql.Field{
					Type: graphql.String,
				},
				"district_id": &graphql.Field{
					Type: graphql.Int,
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
