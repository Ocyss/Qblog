package service

import (
	"github.com/Ocyss/Qblog/cmd/article/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/article"
	"github.com/Ocyss/Qblog/pkg/conv"
)

func (s *ArticleService) Create(request *article.ArticleCreateReq) (resp *article.ArticleCreateResp, err error) {
	resp = new(article.ArticleCreateResp)
	data := db.Article{
		UserId:     uint(request.Userid),
		Title:      request.Data.Title,
		Introduce:  request.Data.Introduce,
		Content:    request.Data.Content,
		Image:      request.Data.Image,
		CategoryID: conv.ToUint(request.Data.Category),
		Show:       request.Data.Show,
	}
	err = db.CreateArticle(s.ctx, &data, request.Data.Tags)
	if err != nil {
		return
	}
	resp.Aid = int32(data.ID)
	resp.Uri = data.Uri
	return
}
