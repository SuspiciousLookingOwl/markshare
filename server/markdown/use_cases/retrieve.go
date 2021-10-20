package markdownUseCases

import (
	"context"

	md "github.com/suspiciouslookingowl/markshare/server/markdown/domains"
	mr "github.com/suspiciouslookingowl/markshare/server/markdown/repositories"
)

type RetriveUsecasePayload struct {
	ID string
}

func (u *MarkdownUseCases) Retrieve(ctx context.Context, payload RetriveUsecasePayload) (*md.Markdown, error) {
	markdowns, _ := u.mdRepo.GetAll(mr.GetAllOptions{
		IDs: []string{payload.ID},
	})

	return &(*markdowns)[0], nil
}
