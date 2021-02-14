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
				//? ============================== Provinces =======================================
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
				//? ============================== END of Provinces =======================================

				//? ============================== Districts =======================================
				"create_district": &graphql.Field{
					Type:        fieldtypes.DistrictType(),
					Description: "Create new district",
					Args:        fieldargs.CreateDistrictArgs(),
					Resolve:     resolvers.CreateDistrictResolver,
				},
				"update_district": &graphql.Field{
					Type:        fieldtypes.DistrictType(),
					Description: "Update district",
					Args:        fieldargs.UpdateDistrictArgs(),
					Resolve:     resolvers.UpdateDistrictResolver,
				},
				"delete_district": &graphql.Field{
					Type:        fieldtypes.DistrictType(),
					Description: "Delete district",
					Args:        fieldargs.DeleteDistrictArgs(),
					Resolve:     resolvers.DeleteDistrictResolver,
				},
				//? ============================== END of Districts =======================================
			},
		},
	)
}
