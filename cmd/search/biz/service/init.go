package service

import (
	"context"
)

type SearchService struct {
	ctx context.Context
}

// NewSearchService new SearchService
func NewSearchService(ctx context.Context) *SearchService {
	return &SearchService{ctx: ctx}
}
