package types

import "github.com/graphql-go/graphql"

var MarkdownConnection = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MarkdownConnection",
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
		},
	},
)
