package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	"github.com/suspiciouslookingowl/markshare/server/api/types"
	markdownDomains "github.com/suspiciouslookingowl/markshare/server/markdown/domains"
	userUseCases "github.com/suspiciouslookingowl/markshare/server/user/use_cases"
)

func NewUserField(app *app.App) *graphql.Field {
	return &graphql.Field{
		Type:        types.User,
		Description: "Get a user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			id, isId := rp.Args["id"].(string)
			md, fromMd := rp.Source.(markdownDomains.Markdown)

			if fromMd && !isId {
				id = md.UserID
			}

			user, _ := app.UserUseCases.Retrieve(rp.Context, userUseCases.RetriveUsecasePayload{
				ID: id,
			})

			return user, nil
		},
	}
}
