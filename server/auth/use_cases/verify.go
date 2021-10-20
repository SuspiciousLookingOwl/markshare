package usecases

func (u *AuthUseCases) Verify(signed string) (string, error) {
	subject, err := u.jwtProvider.Verify(signed)
	if err != nil {
		return "", err
	}

	return subject, err
}
