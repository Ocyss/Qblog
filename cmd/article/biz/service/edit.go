package service

import (
	"github.com/Ocyss/Qblog/cmd/article/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/article"
	"github.com/Ocyss/Qblog/kitex_gen/common"
	"github.com/Ocyss/Qblog/pkg/conv"
)

func (s *ArticleService) Edit(request *article.ArticleEditReq) (resp *common.EmptyStruct, err error) {
	data := db.Article{
		Title:      request.Data.Title,
		Introduce:  request.Data.Introduce,
		Content:    request.Data.Content,
		Image:      request.Data.Image,
		CategoryID: conv.ToUint(request.Data.Category),
		Show:       request.Data.Show,
	}
	err = db.UpdateArticle(s.ctx, uint(request.Aid), request.Uri, data, request.Data.Tags)
	return
}
