package service

import (
	"context"
	"testing"

	"github.com/Ocyss/Qblog/kitex_gen/user"
)

func TestUserService_Login(t *testing.T) {
	s := NewUserService(context.Background())

	request := &user.UserLoginReq{}
	resp, err := s.Login(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test
}
