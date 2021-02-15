package fieldtypes

import "github.com/graphql-go/graphql"

// PersonType type
func PersonType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Person",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"nip": &graphql.Field{
					Type: graphql.String,
				},
				"full_name": &graphql.Field{
					Type: graphql.String,
				},
				"first_name": &graphql.Field{
					Type: graphql.String,
				},
				"last_name": &graphql.Field{
					Type: graphql.String,
				},
				"birth_date": &graphql.Field{
					Type: graphql.String,
				},
				"birth_place": &graphql.Field{
					Type: graphql.String,
				},
				"photo_profile_url": &graphql.Field{
					Type: graphql.String,
				},
				"gender": &graphql.Field{
					Type: graphql.String,
				},
				"zone_location": &graphql.Field{
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
