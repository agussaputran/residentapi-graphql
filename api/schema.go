package api

import "github.com/graphql-go/graphql"

// Schema global var
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    QueryType(),
		Mutation: MutationType(),
	},
)
