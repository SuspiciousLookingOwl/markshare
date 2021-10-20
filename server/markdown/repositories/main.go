package markdownRepositories

import "github.com/doug-martin/goqu"

func NewMarkdownRepository(db *goqu.Database) *MarkdownRepository {
	return &MarkdownRepository{
		db: db,
	}
}

type MarkdownRepository struct {
	db *goqu.Database
}
