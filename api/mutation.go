package api

import (
	"resident-graphql/fieldargs"
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
					Args:        fieldargs.CreateProvinceArgs(),
					Resolve:     resolvers.CreateProvinceResolver,
				},
				"update_province": &graphql.Field{
					Type:        fieldtypes.ProvinceType(),
					Description: "Update province",
					Args:        fieldargs.UpdateProvinceArgs(),
					Resolve:     resolvers.UpdateProvinceResolver,
				},
				"delete_province": &graphql.Field{
					Type:        fieldtypes.ProvinceType(),
					Description: "Delete province",
					Args:        fieldargs.DeleteProvinceArgs(),
					Resolve:     resolvers.DeleteProvinceResolver,
				},
			},
		},
	)
}
