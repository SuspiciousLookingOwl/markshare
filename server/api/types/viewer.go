package types

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	markdownUseCases "github.com/suspiciouslookingowl/markshare/server/markdown/use_cases"
	userDomains "github.com/suspiciouslookingowl/markshare/server/user/domains"
)

var Viewer = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Viewer",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"profilePictureURL": &graphql.Field{
				Type: graphql.String,
			},
			"markdowns": &graphql.Field{
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
				Type: graphql.NewList(MarkdownConnection),
			},
		},
	},
)

func NewViewerType(app *app.App) *graphql.Object {
	Viewer.Fields()["markdowns"].Resolve = func(rp graphql.ResolveParams) (interface{}, error) {
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

	return Viewer
}
