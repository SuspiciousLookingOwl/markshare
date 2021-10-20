package types

import "github.com/graphql-go/graphql"

var Markdown = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Markdown",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"userId": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: User,
			},
		},
	},
)
