package service

import (
	article "github.com/Ocyss/Qblog/kitex_gen/article"
	common "github.com/Ocyss/Qblog/kitex_gen/common"
)

func (s *ArticleService) Delete(request *article.ArticleDeleteReq) (resp *common.EmptyStruct, err error) {
	resp = new(common.EmptyStruct)
	// TODO ArticleService.Delete Finish your business logic.

	return
}
