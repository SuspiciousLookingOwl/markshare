package userUseCases

import (
	user "github.com/suspiciouslookingowl/markshare/server/user/repositories"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type UserUseCases struct {
	Logger *zap.SugaredLogger

	userRepo *user.UserRepository
}

type params struct {
	fx.In

	Logger *zap.SugaredLogger
	*user.UserRepository
}

func NewUserUseCase(p params) *UserUseCases {
	return &UserUseCases{
		Logger:   p.Logger,
		userRepo: p.UserRepository,
	}
}
