package fieldargs

import "github.com/graphql-go/graphql"

// UserLoginArgs for mutation args
func UserLoginArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}
