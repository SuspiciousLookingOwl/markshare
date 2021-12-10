package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/suspiciouslookingowl/markshare/server/api/app"
	"github.com/suspiciouslookingowl/markshare/server/api/mutations"
	"github.com/suspiciouslookingowl/markshare/server/api/queries"
	"github.com/suspiciouslookingowl/markshare/server/api/types"
	c "github.com/suspiciouslookingowl/markshare/server/common"
	"github.com/suspiciouslookingowl/markshare/server/config"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func newSchema(app *app.App) *graphql.Schema {
	schema, _ := graphql.NewSchema(
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
				types.NewMarkdownType(app),
				types.MarkdownConnection,
				types.NewUserType(app),
				types.UserConnection,
				types.NewViewerType(app),
			},
		},
	)
	return &schema
}

func StartServer(app *app.App, env *config.Env) {
	defer app.Logger.Sync()
	app.Logger.Info("Starting server")

	schema := *newSchema(app)

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

		app.Logger.Debugw("New Request",
			"method", req.Method,
			"url", req.URL.Path,
			"body", p,
			"headers", req.Header,
		)

		ctx := req.Context()

		authorization := strings.Split(req.Header.Get("Authorization"), " ")
		if len(authorization) == 2 {
			id, err := app.AuthUseCases.Verify(authorization[1])
			if err != nil {
				w.WriteHeader(401)
				return
			}
			ctx = context.WithValue(req.Context(), c.UserId, id)
		}

		result := graphql.Do(graphql.Params{
			Context: ctx,
			Schema:  schema,

			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		})

		if err := json.NewEncoder(w).Encode(result); err != nil {
			app.Logger.Error("could not write result to response: %s", err)
		}
	})

	app.Logger.Info("Server is running on port " + env.Port)
	http.ListenAndServe(":"+env.Port, nil)
}
