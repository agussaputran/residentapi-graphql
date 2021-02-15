package api

import (
	"resident-graphql/fieldtypes"
	"resident-graphql/resolvers"

	"github.com/graphql-go/graphql"
)

// QueryType func
func QueryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"read_province": &graphql.Field{
					Type:        graphql.NewList(fieldtypes.ProvinceType()),
					Description: "Get Province List",
					Resolve:     resolvers.ReadProvinceResolver,
				},
				"read_district": &graphql.Field{
					Type:        graphql.NewList(fieldtypes.DistrictType()),
					Description: "Get District List",
					Resolve:     resolvers.ReadDistrictResolver,
				},
				"read_sub_district": &graphql.Field{
					Type:        graphql.NewList(fieldtypes.Subs()),
					Description: "Get District List",
					Resolve:     resolvers.ReadDistrictResolver,
				},
			},
		},
	)
}
