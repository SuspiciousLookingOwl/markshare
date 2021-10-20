package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	"github.com/suspiciouslookingowl/markshare/server/api/types"
	markdownUseCases "github.com/suspiciouslookingowl/markshare/server/markdown/use_cases"
)

func NewMarkdownField(app *app.App) *graphql.Field {
	return &graphql.Field{
		Type:        types.Markdown,
		Description: "Get a markdown",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(string)

			markdown, _ := app.MarkdownUseCases.Retrieve(p.Context, markdownUseCases.RetriveUsecasePayload{
				ID: id,
			})

			return markdown, nil
		},
	}

}
