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

	markdown, err := markdownDomains.NewMarkdown(&markdownDomains.Markdown{
		ID:      "", // TODO: Generate
		Content: pl.Content,
		UserID:  userID,
	})

	if err != nil {
		return nil, err
	}

	err = u.mdRepo.Persist(*markdown)
	return markdown, err
}
