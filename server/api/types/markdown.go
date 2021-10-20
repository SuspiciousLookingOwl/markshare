package types

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	markdownDomains "github.com/suspiciouslookingowl/markshare/server/markdown/domains"
	userUseCases "github.com/suspiciouslookingowl/markshare/server/user/use_cases"
)

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
				Type: UserConnection,
			},
		},
	},
)

func NewMarkdownType(app *app.App) *graphql.Object {
	fields := Markdown.Fields()

	fields["author"].Resolve = func(rp graphql.ResolveParams) (interface{}, error) {
		markdown, _ := rp.Source.(markdownDomains.Markdown)

		options := userUseCases.RetriveUsecasePayload{
			ID: markdown.UserID,
		}
		user, _ := app.UserUseCases.Retrieve(rp.Context, options)

		return user, nil
	}

	return User
}
