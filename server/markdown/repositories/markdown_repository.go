package markdownRepositories

import (
	"errors"

	"github.com/doug-martin/goqu"
	md "github.com/suspiciouslookingowl/markshare/server/markdown/domains"
)

var GetErrors = struct {
	ProviderNotFound error
}{
	ProviderNotFound: errors.New("provider not found"),
}

func (r *MarkdownRepository) Persist(markdown md.Markdown) error {
	_, err := r.db.From("markdowns").Insert(markdown).Exec()
	return err
}

// Get All
type GetAllOptions struct {
	IDs            []string
	UserIDs        []string
	After          string
	Limit          uint
	OrderBy        string
	OrderDirection string
}

func (r *MarkdownRepository) GetAll(opt GetAllOptions) (*[]md.Markdown, error) {
	qb := r.db.From("markdowns")

	if len(opt.UserIDs) != 0 {
		qb = qb.Where(goqu.I("user_id").In(opt.UserIDs))
	}
	if len(opt.IDs) != 0 {
		qb = qb.Where(goqu.I("id").In(opt.IDs))
	}

	if opt.After != "" {
		qb = qb.Where(goqu.I("id").Gt(opt.After))
	}

	limit := uint(1)
	if opt.Limit > 0 {
		limit = opt.Limit
	}
	qb = qb.Limit(limit)

	if opt.OrderBy != "" {
		orderEx := goqu.I(opt.OrderBy)
		if opt.OrderDirection != "ASC" {
			qb = qb.Order(orderEx.Asc())
		} else {
			qb = qb.Order(orderEx.Desc())
		}
	}

	var markdowns []md.Markdown
	qb.ScanStructs(&markdowns)

	return &markdowns, nil
}
