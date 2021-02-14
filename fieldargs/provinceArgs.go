package fieldargs

import "github.com/graphql-go/graphql"

// CreateProvinceArgs for mutation args
func CreateProvinceArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"province_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

// UpdateProvinceArgs for mutation args
func UpdateProvinceArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"province_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

// DeleteProvinceArgs for mutation args
func DeleteProvinceArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
