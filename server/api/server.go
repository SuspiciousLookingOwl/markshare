package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	"github.com/suspiciouslookingowl/markshare/server/config"

	"github.com/suspiciouslookingowl/markshare/server/api/mutations"
	"github.com/suspiciouslookingowl/markshare/server/api/queries"
	"github.com/suspiciouslookingowl/markshare/server/api/types"
	authUsecase "github.com/suspiciouslookingowl/markshare/server/auth/use_cases"
	markdownUseCases "github.com/suspiciouslookingowl/markshare/server/markdown/use_cases"
	userUseCases "github.com/suspiciouslookingowl/markshare/server/user/use_cases"
	"go.uber.org/fx"
)

type param struct {
	fx.In

	*userUseCases.UserUseCases
	*markdownUseCases.MarkdownUseCases
	*authUsecase.AuthUseCases
}

func NewApp(p param) *app.App {

	app := &app.App{
		UserUseCases:     *p.UserUseCases,
		MarkdownUseCases: *p.MarkdownUseCases,
		AuthUseCases:     *p.AuthUseCases,
	}

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"user":      queries.NewUserField(app),
					"viewer":    queries.NewViewerField(app),
					"markdown":  queries.NewMarkdownField(app),
					"markdowns": queries.NewMarkdownsField(app),
				},
			}),
			Mutation: graphql.NewObject(graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"login":          mutations.NewLoginMutation(app),
					"createMarkdown": mutations.NewCreateMarkdownMutation(app),
				},
			}),
			Types: []graphql.Type{
				types.Markdown,
				types.MarkdownConnection,
				types.NewUserType(app),
				types.UserConnection,
				types.NewViewerType(app),
			},
		},
	)

	app.Schema = &schema

	return app
}

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func StartServer(a *app.App, env *config.Env) {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {

		// CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
		if req.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}

		var p postData
		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			w.WriteHeader(400)
			return
		}

		ctx := req.Context()

		authorization := strings.Split(req.Header.Get("Authorization"), " ")
		if len(authorization) == 2 {
			id, err := a.AuthUseCases.Verify(authorization[1])
			if err != nil {
				w.WriteHeader(401)
				return
			}
			ctx = context.WithValue(req.Context(), app.UserId, id)
		}

		result := graphql.Do(graphql.Params{
			Context: ctx,
			Schema:  *a.Schema,

			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		})

		if err := json.NewEncoder(w).Encode(result); err != nil {
			fmt.Printf("could not write result to response: %s", err)
		}
	})

	fmt.Println("Now server is running on port ", env.Port)
	http.ListenAndServe(":"+env.Port, nil)
}
