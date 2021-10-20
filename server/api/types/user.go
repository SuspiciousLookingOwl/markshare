package types

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	markdownUseCases "github.com/suspiciouslookingowl/markshare/server/markdown/use_cases"
	userDomains "github.com/suspiciouslookingowl/markshare/server/user/domains"
)

var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
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
			"markdowns": &graphql.Field{
				Type: graphql.NewList(MarkdownConnection),
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 10,
					},
					"after": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
			},
		},
	},
)

func NewUserType(app *app.App) *graphql.Object {
	fields := User.Fields()

	fields["markdowns"].Resolve = func(rp graphql.ResolveParams) (interface{}, error) {
		userID := rp.Source.(*userDomains.User).ID
		limitInt, _ := rp.Args["limit"].(int)
		limit := uint(limitInt)
		after, _ := rp.Args["after"].(string)

		options := markdownUseCases.RetrieveAllUsecasePayload{
			UserIDs: []string{userID},
			Limit:   limit,
			After:   after,
		}

		markdowns, _ := app.MarkdownUseCases.RetrieveAll(rp.Context, options)

		return markdowns, nil
	}

	return User
}
