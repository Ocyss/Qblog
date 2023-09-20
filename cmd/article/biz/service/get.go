package service

import (
	"github.com/Ocyss/Qblog/cmd/article/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/article"
	"github.com/Ocyss/Qblog/pkg/conv"
)

func (s *ArticleService) Get(request *article.ArticleReq) (resp *article.ArticleResp, err error) {
	resp = new(article.ArticleResp)
	data, err := db.GetArticle(s.ctx, uint(request.Aid), request.Uri)
	resp.Article = &article.Article{
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
	return
}

func ToStringsTags(s []string) (res []db.Tag) {
	if n := len(s); n > 0 {
		res = make([]db.Tag, n)
		for i, tag := range s {
			res[i].Name = tag
		}
	}
	return
}

func ToTagsStrings(s []db.Tag) (res []string) {
	if n := len(s); n > 0 {
		res = make([]string, n)
		for i, tag := range s {
			res[i] = tag.Name
		}
	}
	return
}
