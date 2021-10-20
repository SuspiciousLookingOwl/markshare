package userrepo

import (
	"github.com/doug-martin/goqu"
	ud "github.com/suspiciouslookingowl/markshare/server/user/domains"
)

type GetAllOptions struct {
	IDs    []string
	Emails []string
	Limit  uint
}

func (r *UserRepository) GetAll(opt GetAllOptions) (*[]ud.User, error) {
	qb := r.db.From("users")

	if len(opt.Emails) > 0 {
		qb = qb.Where(goqu.I("email").In(opt.Emails))
	}
	if len(opt.IDs) > 0 {
		qb = qb.Where(goqu.I("id").In(opt.IDs))
	}

	limit := uint(1)
	if opt.Limit > 0 {
		limit = opt.Limit
	}
	qb = qb.Limit(limit)

	var users []ud.User
	qb.ScanStructs(&users)

	return &users, nil
}

func (r *UserRepository) Persist(user ud.User) error {
	_, err := r.db.From("users").Insert(user).Exec()
	return err
}
