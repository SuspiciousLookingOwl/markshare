package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	"github.com/suspiciouslookingowl/markshare/server/api/types"
	c "github.com/suspiciouslookingowl/markshare/server/common"
	userUseCases "github.com/suspiciouslookingowl/markshare/server/user/use_cases"
)

func NewViewerField(app *app.App) *graphql.Field {
	return &graphql.Field{
		Type:        types.Viewer,
		Description: "Get viewer data",
		Args:        graphql.FieldConfigArgument{},
		Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
			id, _ := rp.Context.Value(c.UserId).(string)

			if id == "" {
				return nil, nil
			}

			user, _ := app.UserUseCases.Retrieve(rp.Context, userUseCases.RetriveUsecasePayload{
				ID: id,
			})

			return user, nil
		},
	}
}
