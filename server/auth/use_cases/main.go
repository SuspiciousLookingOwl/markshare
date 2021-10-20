package usecases

import (
	"github.com/suspiciouslookingowl/markshare/server/auth/providers"
	user "github.com/suspiciouslookingowl/markshare/server/user/repositories"
	"go.uber.org/fx"
)

type AuthUseCases struct {
	userRepo     *user.UserRepository
	authProvider *providers.GitHubAuthProvider
	jwtProvider  *providers.JWTProvider
}

type param struct {
	fx.In

	*user.UserRepository
	*providers.GitHubAuthProvider
	*providers.JWTProvider
}

func NewAuthUseCase(p param) *AuthUseCases {
	return &AuthUseCases{
		userRepo:     p.UserRepository,
		authProvider: p.GitHubAuthProvider,
		jwtProvider:  p.JWTProvider,
	}
}
