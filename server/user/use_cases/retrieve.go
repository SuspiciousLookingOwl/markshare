package userUseCases

import (
	"context"

	ud "github.com/suspiciouslookingowl/markshare/server/user/domains"
	ur "github.com/suspiciouslookingowl/markshare/server/user/repositories"
)

type RetriveUsecasePayload struct {
	ID    string
	Email string
}

func (u *UserUseCases) Retrieve(ctx context.Context, pl RetriveUsecasePayload) (*ud.User, error) {
	var options ur.GetAllOptions

	if pl.Email != "" {
		options.Emails = []string{pl.Email}
	}
	if pl.ID != "" {
		options.IDs = []string{pl.ID}
	}

	users, _ := u.userRepo.GetAll(options)

	var user *ud.User
	if len(*users) > 0 {
		user = &(*users)[0]
	}

	userID, _ := ctx.Value("user_id").(string)

	if user != nil && userID != user.ID {
		user.Email = ""
	}

	return user, nil
}
