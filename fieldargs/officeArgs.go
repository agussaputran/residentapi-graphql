package fieldargs

import "github.com/graphql-go/graphql"

// CreateOfficeArgs for mutation args
func CreateOfficeArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"office_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"sub_district_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}

// UpdateOfficeArgs for mutation args
func UpdateOfficeArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"office_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

// DeleteOfficeArgs for mutation args
func DeleteOfficeArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
