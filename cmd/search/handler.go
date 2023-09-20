package main

import (
	"context"

	"github.com/Ocyss/Qblog/cmd/search/biz/service"
	"github.com/Ocyss/Qblog/kitex_gen/search"
)

// SearchServiceImpl implements the last service interface defined in the IDL.
type SearchServiceImpl struct{}

// FindArticle implements the SearchServiceImpl interface.
func (s *SearchServiceImpl) FindArticle(ctx context.Context, request *search.SearchArticleReq) (resp *search.SearchArticleResp, err error) {
	searchService := service.NewSearchService(ctx)
	resp, err = searchService.FindArticle(request)
	return
}
