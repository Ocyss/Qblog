package service

import (
	"github.com/Ocyss/Qblog/cmd/article/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/article"
)

func (s *ArticleService) GetIdUri(request *article.ArticleIdUriReq) (resp *article.ArticleIdUriResp, err error) {
	id, uri, err := db.GetIdUri(s.ctx, uint(request.Id), request.Uri)
	if err != nil {
		return
	}
	resp = &article.ArticleIdUriResp{Uri: uri, Id: int64(id)}
	return
}
