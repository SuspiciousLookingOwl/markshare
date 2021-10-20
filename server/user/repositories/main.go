package userrepo

import "github.com/doug-martin/goqu"

func NewUserRepository(db *goqu.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *goqu.Database
}
