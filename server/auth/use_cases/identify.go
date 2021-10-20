package usecases

import (
	"errors"

	userDomains "github.com/suspiciouslookingowl/markshare/server/user/domains"
	userrepo "github.com/suspiciouslookingowl/markshare/server/user/repositories"
)

type IdentifyUsecasePayload struct {
	Credential string
	Provider   string
}

func (u *AuthUseCases) Identify(pl IdentifyUsecasePayload) (string, error) {
	if pl.Provider == "github" {
		accessToken, err := u.authProvider.GetAccessToken(pl.Credential)
		if err != nil {
			return "", err
		}

		primaryEmail, err := u.authProvider.GetUserPrimaryEmail(accessToken)
		if err != nil {
			return "", err
		}

		users, _ := u.userRepo.GetAll(userrepo.GetAllOptions{
			Emails: []string{primaryEmail},
		})

		user := &(*users)[0]

		// if user doesn't exists, create new user
		if user == nil {
			ghUser, err := u.authProvider.GetUser(accessToken)
			if err != nil {
				return "", err
			}

			user = &userDomains.User{
				ID:                ghUser.Login,
				Email:             primaryEmail,
				Name:              ghUser.Name,
				ProfilePictureURL: ghUser.AvatarURL,
			}

			u.userRepo.Persist(*user)
		}

		// create jwt
		signed, err := u.jwtProvider.Sign(user.ID)

		if err != nil {
			return "", err
		}

		return signed, nil
	}

	return "", errors.New("invalid provider")

}
