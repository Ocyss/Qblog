package service

import (
	"github.com/Ocyss/Qblog/cmd/article/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/article"
)

func (s *ArticleService) CheckIdUri(request *article.ArticleIdUriReq) (err error) {
	_, _, err = db.GetIdUri(s.ctx, uint(request.Id), request.Uri)
	return
}
