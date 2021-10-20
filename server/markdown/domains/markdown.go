package markdownDomains

import "github.com/go-playground/validator/v10"

type Markdown struct {
	ID      string `json:"id"  db:"id" validate:"required"`
	Content string `json:"content" db:"content" validate:"required,min=3,max=8096"`
	UserID  string `json:"userId" db:"user_id" validate:"required"`
}

func NewMarkdown(markdown *Markdown) (*Markdown, error) {
	v := validator.New()
	err := v.Struct(markdown)

	if err != nil {
		return nil, err
	}

	return markdown, nil
}
