package app

import (
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

func NewApp(p param) *App {

	app := &App{
		UserUseCases:     *p.UserUseCases,
		MarkdownUseCases: *p.MarkdownUseCases,
		AuthUseCases:     *p.AuthUseCases,
	}

	return app
}
