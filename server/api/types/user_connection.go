package types

import "github.com/graphql-go/graphql"

var UserConnection = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserConnection",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"profilePictureURL": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
