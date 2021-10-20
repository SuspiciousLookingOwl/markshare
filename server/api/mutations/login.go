package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	authUseCases "github.com/suspiciouslookingowl/markshare/server/auth/use_cases"
)

type LoginResponse struct {
	Token string `json:"token" db:"id"`
}

func NewLoginMutation(app *app.App) *graphql.Field {
	return &graphql.Field{
		Name: "Login",
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Credential",
			Fields: graphql.Fields{
				"token": &graphql.Field{
					Type: graphql.String,
				},
			},
		}),
		Description: "Login",
		Args: graphql.FieldConfigArgument{
			"credential": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"provider": &graphql.ArgumentConfig{
				Type:         graphql.String,
				Description:  "OAuth Provider (always github)",
				DefaultValue: "github",
			},
		},
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			credential := rp.Args["credential"].(string)
			provider := rp.Args["provider"].(string)

			token, err := app.AuthUseCases.Identify(authUseCases.IdentifyUsecasePayload{
				Credential: credential,
				Provider:   provider,
			})

			if err != nil {
				return nil, err
			}

			return LoginResponse{
				Token: token,
			}, nil
		},
	}
}
