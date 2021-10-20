package app

import (
	"github.com/graphql-go/graphql"
	authUsecase "github.com/suspiciouslookingowl/markshare/server/auth/use_cases"
	markdownUseCases "github.com/suspiciouslookingowl/markshare/server/markdown/use_cases"
	userUseCases "github.com/suspiciouslookingowl/markshare/server/user/use_cases"
)

type App struct {
	userUseCases.UserUseCases
	markdownUseCases.MarkdownUseCases
	authUsecase.AuthUseCases

	Schema *graphql.Schema
}

type UserIdType string

const UserId UserIdType = "user_id"
