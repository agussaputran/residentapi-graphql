package fieldtypes

import "github.com/graphql-go/graphql"

// DistrictType type
func DistrictType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "District",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"district_name": &graphql.Field{
					Type: graphql.String,
				},
				"province_id": &graphql.Field{
					Type: graphql.Int,
				},
				"province_name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
