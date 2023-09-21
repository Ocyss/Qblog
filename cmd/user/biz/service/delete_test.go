package service

import (
	"context"
	"testing"

	user "github.com/Ocyss/Qblog/kitex_gen/user"
)

func TestUserService_Delete(t *testing.T) {
	s := NewUserService(context.Background())

	request := &user.UserDeleteReq{}
	resp, err := s.Delete(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// TEST: edit your unit test
}
