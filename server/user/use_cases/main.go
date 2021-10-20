package userUseCases

import (
	user "github.com/suspiciouslookingowl/markshare/server/user/repositories"
)

type UserUseCases struct {
	userRepo *user.UserRepository
}

func NewUserUseCase(userRepo *user.UserRepository) *UserUseCases {
	return &UserUseCases{
		userRepo: userRepo,
	}
}
