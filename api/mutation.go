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

				//? ============================== Sub Districts =======================================
				"create_sub_district": &graphql.Field{
					Type:        fieldtypes.SubDistrictType(),
					Description: "Create new sub district",
					Args:        fieldargs.CreateSubDistrictArgs(),
					Resolve:     resolvers.CreateSubDistrictResolver,
				},
				"update_sub_district": &graphql.Field{
					Type:        fieldtypes.SubDistrictType(),
					Description: "Update sub district",
					Args:        fieldargs.UpdateSubDistrictArgs(),
					Resolve:     resolvers.UpdateSubDistrictResolver,
				},
				"delete_sub_district": &graphql.Field{
					Type:        fieldtypes.SubDistrictType(),
					Description: "Delete district",
					Args:        fieldargs.DeleteSubDistrictArgs(),
					Resolve:     resolvers.DeleteSubDistrictResolver,
				},
				//? ============================== END of Sub Districts =======================================

				//? ============================== Sub Persons =======================================
				"create_person": &graphql.Field{
					Type:        fieldtypes.PersonType(),
					Description: "Create new person",
					Args:        fieldargs.CreatePersonArgs(),
					Resolve:     resolvers.CreatePersonResolver,
				},
				"update_person": &graphql.Field{
					Type:        fieldtypes.PersonType(),
					Description: "Update person",
					Args:        fieldargs.UpdatePersonArgs(),
					Resolve:     resolvers.UpdatePersonResolver,
				},
				"delete_person": &graphql.Field{
					Type:        fieldtypes.PersonType(),
					Description: "Delete person",
					Args:        fieldargs.DeletePersonArgs(),
					Resolve:     resolvers.DeletePersonResolver,
				},
				//? ============================== END of Persons =======================================

				//? ============================== Sub Persons =======================================
				"create_office": &graphql.Field{
					Type:        fieldtypes.OfficeType(),
					Description: "Create new office",
					Args:        fieldargs.CreateOfficeArgs(),
					Resolve:     resolvers.CreateOfficeResolver,
				},
				//? ============================== END of Persons =======================================
			},
		},
	)
}
