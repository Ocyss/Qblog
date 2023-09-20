package service

import (
	"github.com/Ocyss/Qblog/cmd/article/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/article"
	"github.com/Ocyss/Qblog/pkg/conv"
)

func (s *ArticleService) GetList(request *article.ArticlesReq) (resp *article.ArticlesResp, err error) {
	resp = new(article.ArticlesResp)
	datas, err := db.GetArticleList(s.ctx, int(request.PageSize), int(request.PageNum), int(request.Category), int(request.Tag))
	if err != nil {
		return
	}
	articles := make([]*article.Article, len(datas))
	for i, data := range datas {
		articles[i] = &article.Article{
			Aid:        int64(data.ID),
			Uri:        data.Uri,
			Title:      data.Title,
			Introduce:  data.Introduce,
			Content:    data.Content,
			Image:      data.Image,
			Tags:       ToTagsStrings(data.Tags),
			Category:   conv.ToInt64(data.CategoryID),
			Show:       data.Show,
			Uv:         0,
			Pv:         0,
			CreateTime: data.CreatedAt.Unix(),
			UpdateTime: data.UpdatedAt.Unix(),
		}
	}
	resp.Articles = articles
	return
}
