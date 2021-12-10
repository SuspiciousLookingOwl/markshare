package app

import (
	authUsecase "github.com/suspiciouslookingowl/markshare/server/auth/use_cases"
	markdownUseCases "github.com/suspiciouslookingowl/markshare/server/markdown/use_cases"
	userUseCases "github.com/suspiciouslookingowl/markshare/server/user/use_cases"
	"go.uber.org/zap"
)

type App struct {
	Logger zap.SugaredLogger

	userUseCases.UserUseCases
	markdownUseCases.MarkdownUseCases
	authUsecase.AuthUseCases
}
