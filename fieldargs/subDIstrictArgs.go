package fieldargs

import "github.com/graphql-go/graphql"

// CreateSubDistrictArgs for mutation args
func CreateSubDistrictArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"district_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"district_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}

// UpdateSubDistrictArgs for mutation args
func UpdateSubDistrictArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"sub_district_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

// DeleteSubDistrictArgs for mutation args
func DeleteSubDistrictArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
