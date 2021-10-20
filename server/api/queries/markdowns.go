package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	"github.com/suspiciouslookingowl/markshare/server/api/types"
	markdownUseCases "github.com/suspiciouslookingowl/markshare/server/markdown/use_cases"
)

func NewMarkdownsField(app *app.App) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.Markdown),
		Description: "List of markdowns",
		Args: graphql.FieldConfigArgument{
			"authorId": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"limit": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: 10,
			},
			"after": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
		},
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			userID, _ := rp.Args["authorId"].(string)
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
		},
	}
}
