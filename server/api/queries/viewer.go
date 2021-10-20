package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	"github.com/suspiciouslookingowl/markshare/server/api/types"
	userUseCases "github.com/suspiciouslookingowl/markshare/server/user/use_cases"
)

func NewViewerField(a *app.App) *graphql.Field {
	return &graphql.Field{
		Type:        types.Viewer,
		Description: "Get viewer data",
		Args:        graphql.FieldConfigArgument{},
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			id, _ := rp.Context.Value(app.UserId).(string)

			if id == "" {
				return nil, nil
			}

			user, _ := a.UserUseCases.Retrieve(rp.Context, userUseCases.RetriveUsecasePayload{
				ID: id,
			})

			return user, nil
		},
	}
}
