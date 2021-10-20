package markdownUseCases

import (
	markdownRepository "github.com/suspiciouslookingowl/markshare/server/markdown/repositories"
)

type MarkdownUseCases struct {
	mdRepo *markdownRepository.MarkdownRepository
}

func NewMarkdownUseCases(mdRepo *markdownRepository.MarkdownRepository) *MarkdownUseCases {
	return &MarkdownUseCases{
		mdRepo: mdRepo,
	}
}
