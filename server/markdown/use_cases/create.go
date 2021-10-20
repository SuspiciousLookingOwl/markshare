package markdownUseCases

import (
	"context"
	"errors"

	c "github.com/suspiciouslookingowl/markshare/server/common"
	markdownDomains "github.com/suspiciouslookingowl/markshare/server/markdown/domains"
)

type CreateUsecasePayload struct {
	Content string
}

func (u *MarkdownUseCases) Create(ctx context.Context, pl CreateUsecasePayload) (*markdownDomains.Markdown, error) {
	userID := ctx.Value(c.UserId).(string)
	if userID == "" {
		return nil, errors.New("user not found")
	}

	markdown := &markdownDomains.Markdown{
		ID:      "random234",
		Content: pl.Content,
		UserID:  ctx.Value("user_id").(string),
	}

	err := u.mdRepo.Persist(*markdown)
	return markdown, err
}
