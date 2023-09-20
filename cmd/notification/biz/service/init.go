package service

import (
	"context"
)

type NotificationService struct {
	ctx context.Context
}

// NewNotificationService new NotificationService
func NewNotificationService(ctx context.Context) *NotificationService {
	return &NotificationService{ctx: ctx}
}
