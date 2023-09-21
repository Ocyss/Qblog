package service

import (
	"context"
	"testing"

	notification "github.com/Ocyss/Qblog/kitex_gen/notification"
)

func TestNotificationService_Send(t *testing.T) {
	s := NewNotificationService(context.Background())

	request := &notification.NotificationSendReq{}
	resp, err := s.Send(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// TEST: edit your unit test
}
