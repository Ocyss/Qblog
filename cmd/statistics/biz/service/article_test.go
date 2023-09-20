package service

import (
	"context"
	statistics "github.com/Ocyss/Qblog/kitex_gen/statistics"
	"testing"
)

func TestStatisticsService_Article(t *testing.T) {
	s := NewStatisticsService(context.Background())

	request := &statistics.StatisticsArticleReq{}
	resp, err := s.Article(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
