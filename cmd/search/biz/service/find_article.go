package service

import (
	"github.com/Ocyss/Qblog/kitex_gen/search"
)

func (s *SearchService) FindArticle(request *search.SearchArticleReq) (resp *search.SearchArticleResp, err error) {
	resp = new(search.SearchArticleResp)
	// TODO SearchService.FindArticle Finish your business logic.

	return
}
