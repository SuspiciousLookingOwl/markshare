package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	"github.com/suspiciouslookingowl/markshare/server/api/types"
	markdownUseCases "github.com/suspiciouslookingowl/markshare/server/markdown/use_cases"
)

func NewCreateMarkdownMutation(app *app.App) *graphql.Field {
	return &graphql.Field{
		Type:        types.Markdown,
		Description: "Create Markdown",
		Args: graphql.FieldConfigArgument{
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			content := rp.Args["content"].(string)

			markdown, err := app.MarkdownUseCases.Create(rp.Context, markdownUseCases.CreateUsecasePayload{
				Content: content,
			})

			if err != nil {
				return nil, err
			}

			return markdown, nil
		},
	}
}
