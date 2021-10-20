package main

import (
	"database/sql"
	"os"

	"github.com/doug-martin/goqu"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"

	server "github.com/suspiciouslookingowl/markshare/server/api"
	app "github.com/suspiciouslookingowl/markshare/server/api/app"
	authProviders "github.com/suspiciouslookingowl/markshare/server/auth/providers"
	authusecases "github.com/suspiciouslookingowl/markshare/server/auth/use_cases"
	"github.com/suspiciouslookingowl/markshare/server/config"
	markdownRepositories "github.com/suspiciouslookingowl/markshare/server/markdown/repositories"
	markdownUseCases "github.com/suspiciouslookingowl/markshare/server/markdown/use_cases"
	userRepo "github.com/suspiciouslookingowl/markshare/server/user/repositories"
	userUseCases "github.com/suspiciouslookingowl/markshare/server/user/use_cases"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load(".env")

	fx.New(
		// Config
		fx.Provide(config.NewConf),

		// DB
		fx.Provide(func() *goqu.Database {
			db, err := sql.Open("pgx", os.Getenv("POSTGRES_URL"))
			if err != nil {
				panic(err)
			}

			return &goqu.Database{
				Dialect: "postgres",
				Db:      db,
			}
		}),

		// App
		fx.Provide(app.NewApp),

		// Usecases
		fx.Provide(
			userUseCases.NewUserUseCase,
			authusecases.NewAuthUseCase,
			markdownUseCases.NewMarkdownUseCases,
		),

		//Repositories
		fx.Provide(
			userRepo.NewUserRepository,
			markdownRepositories.NewMarkdownRepository,
		),

		// Providers
		fx.Provide(
			authProviders.NewGitHubAuthProvider,
			authProviders.NewJWTProvider,
		),

		// Start
		fx.Invoke(server.StartServer),
	).Run()

}
