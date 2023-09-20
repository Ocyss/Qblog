package service

import (
	article "github.com/Ocyss/Qblog/kitex_gen/article"
	common "github.com/Ocyss/Qblog/kitex_gen/common"
)

func (s *ArticleService) Edit(request *article.ArticleEditReq) (resp *common.EmptyStruct, err error) {
	resp = new(common.EmptyStruct)
	// TODO ArticleService.Edit Finish your business logic.

	return
}
