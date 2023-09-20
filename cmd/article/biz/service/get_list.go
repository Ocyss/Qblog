package service

import (
	article "github.com/Ocyss/Qblog/kitex_gen/article"
)

func (s *ArticleService) GetList(request *article.ArticlesReq) (resp *article.ArticlesResp, err error) {
	resp = new(article.ArticlesResp)
	// TODO ArticleService.GetList Finish your business logic.

	return
}
