package service

import (
	article "github.com/Ocyss/Qblog/kitex_gen/article"
)

func (s *ArticleService) Create(request *article.ArticleCreateReq) (resp *article.ArticleCreateResp, err error) {
	resp = new(article.ArticleCreateResp)
	// TODO ArticleService.Create Finish your business logic.

	return
}
