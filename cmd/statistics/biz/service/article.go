package service

import (
	"errors"

	"github.com/Ocyss/Qblog/cmd/statistics/biz/dal/redis"

	"github.com/Ocyss/Qblog/cmd/statistics/biz/rpc"
	"github.com/Ocyss/Qblog/kitex_gen/article"
	"github.com/Ocyss/Qblog/kitex_gen/statistics"
)

func (s *StatisticsService) Article(request *statistics.StatisticsArticleReq) (resp *statistics.StatisticsArticleResp, err error) {
	if request.Id == 0 {
		if request.Uri == "" {
			err = errors.New("obtaining statistical data must have an id or uri")
			return
		}
		articleIdUriResp, err := rpc.Article.GetIdUri(s.ctx, &article.ArticleIdUriReq{Uri: request.Uri})
		if err != nil {
			return
		}
		request.Id = articleIdUriResp.Id
	}
	uv, pv, err := redis.GetUvPv(s.ctx, uint(request.Id), request.Ua)
	resp = &statistics.StatisticsArticleResp{Pv: int32(pv), Uv: int32(uv)}
	return
}
