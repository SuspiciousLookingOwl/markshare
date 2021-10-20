package config

import "os"

type Env struct {
	GitHubClientID     string
	GitHubClientSecret string
	JWTSecret          string
	PostgresURL        string
	Port               string
}

func NewConf() *Env {
	return &Env{
		PostgresURL:        os.Getenv("DATABASE_URL"),
		GitHubClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		GitHubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		Port:               os.Getenv("PORT"),
	}
}
