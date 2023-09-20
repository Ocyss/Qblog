package service

import (
	article "github.com/Ocyss/Qblog/kitex_gen/article"
)

func (s *ArticleService) GetCategoryList() (resp *article.CategorysResp, err error) {
	resp = new(article.CategorysResp)
	// TODO ArticleService.GetCategoryList Finish your business logic.

	return
}
