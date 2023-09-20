package main

import (
	"context"
	statistics "github.com/Ocyss/Qblog/kitex_gen/statistics"

	"github.com/Ocyss/Qblog/cmd/statistics/biz/service"
)

// StatisticsServiceImpl implements the last service interface defined in the IDL.
type StatisticsServiceImpl struct{}

// Article implements the StatisticsServiceImpl interface.
func (s *StatisticsServiceImpl) Article(ctx context.Context, request *statistics.StatisticsArticleReq) (resp *statistics.StatisticsArticleResp, err error) {
	statisticsService := service.NewStatisticsService(ctx)
	resp, err = statisticsService.Article(request)
	return
}
