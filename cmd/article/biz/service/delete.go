package service

import (
	"github.com/Ocyss/Qblog/cmd/article/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/article"
	"github.com/Ocyss/Qblog/kitex_gen/common"
)

func (s *ArticleService) Delete(request *article.ArticleDeleteReq) (resp *common.EmptyStruct, err error) {
	del := false
	if request.IsDelete != nil {
		del = *request.IsDelete
	}
	err = db.DeleteArticle(s.ctx, uint(request.Aid), request.Uri, del)
	return
}
