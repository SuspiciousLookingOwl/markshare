package markdownUseCases

import (
	"context"

	md "github.com/suspiciouslookingowl/markshare/server/markdown/domains"
	mr "github.com/suspiciouslookingowl/markshare/server/markdown/repositories"
)

type RetrieveAllUsecasePayload struct {
	IDs            []string
	UserIDs        []string
	After          string
	Limit          uint
	OrderBy        string
	OrderDirection string
}

func (u *MarkdownUseCases) RetrieveAll(ctx context.Context, payload RetrieveAllUsecasePayload) (*[]md.Markdown, error) {
	options := mr.GetAllOptions(payload) // probably don't do this lol

	markdown, _ := u.mdRepo.GetAll(options)

	return markdown, nil
}
