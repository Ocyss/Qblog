package service

import (
	article "github.com/Ocyss/Qblog/kitex_gen/article"
)

func (s *ArticleService) Get(request *article.ArticleReq) (resp *article.ArticleResp, err error) {
	resp = new(article.ArticleResp)
	// TODO ArticleService.Get Finish your business logic.

	return
}
