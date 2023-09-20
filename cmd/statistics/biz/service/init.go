package service

import (
	"context"
)

type StatisticsService struct {
	ctx context.Context
}

// NewStatisticsService new StatisticsService
func NewStatisticsService(ctx context.Context) *StatisticsService {
	return &StatisticsService{ctx: ctx}
}
