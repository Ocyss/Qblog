package main

import (
	"context"

	"github.com/Ocyss/Qblog/cmd/notification/biz/service"
	"github.com/Ocyss/Qblog/kitex_gen/common"
	"github.com/Ocyss/Qblog/kitex_gen/notification"
)

// NotificationServiceImpl implements the last service interface defined in the IDL.
type NotificationServiceImpl struct{}

// Send implements the NotificationServiceImpl interface.
func (s *NotificationServiceImpl) Send(ctx context.Context, request *notification.NotificationSendReq) (resp *common.EmptyStruct, err error) {
	notificationService := service.NewNotificationService(ctx)
	resp, err = notificationService.Send(request)
	return
}
