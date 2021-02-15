package fieldargs

import "github.com/graphql-go/graphql"

// CreatePersonArgs for mutation args
func CreatePersonArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"nip": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"full_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"first_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"last_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_date": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_place": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"gender": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"zone_location": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"sub_district_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}

// UpdatePersonArgs for mutation args
func UpdatePersonArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"nip": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"full_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"first_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"last_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_date": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_place": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"gender": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"zone_location": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"sub_district_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}

// DeletePersonArgs for mutation args
func DeletePersonArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
