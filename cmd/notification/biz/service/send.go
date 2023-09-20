package service

import (
	"github.com/Ocyss/Qblog/kitex_gen/common"

	"github.com/Ocyss/Qblog/kitex_gen/notification"
)

func (s *NotificationService) Send(request *notification.NotificationSendReq) (resp *common.EmptyStruct, err error) {
	resp = new(common.EmptyStruct)
	// TODO NotificationService.Send Finish your business logic.

	return
}
