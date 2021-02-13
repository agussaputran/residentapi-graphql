package api

import (
	"resident-graphql/fieldtypes"
	"resident-graphql/resolvers"

	"github.com/graphql-go/graphql"
)

// MutationType func
func MutationType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"create_province": &graphql.Field{
					Type:        fieldtypes.ProvinceType(),
					Description: "Create new province",
					Args:        typesargs.CreateProvinceArgs(),
					Resolve:     resolvers.CreateProvinceResolver,
				},
			},
		},
	)
}
