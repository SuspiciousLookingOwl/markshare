package providers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/suspiciouslookingowl/markshare/server/config"
	"go.uber.org/fx"
)

// Constructor
type JWTProvider struct {
	Secret string
}

type newJWTProviderParam struct {
	fx.In

	*config.Env
}

func NewJWTProvider(p newJWTProviderParam) *JWTProvider {
	return &JWTProvider{
		Secret: p.Env.JWTSecret,
	}
}

func (p *JWTProvider) Sign(subject string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 30 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   subject,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString([]byte(p.Secret))
	return signed, nil
}

func (p *JWTProvider) Verify(signed string) (string, error) {
	token, err := jwt.ParseWithClaims(signed, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(p.Secret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", errors.New("invalid claim")
	}

	return claims.Subject, nil
}
