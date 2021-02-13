package api

import (
	"log"

	"github.com/graphql-go/graphql"
)

// Schema global var
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    QueryType(),
		Mutation: MutationType(),
	},
)

// ExecuteQuery func
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Println(result.Errors)
	}
	return result
}
