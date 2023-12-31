package service

import (
	"context"
	user "github.com/Ocyss/Qblog/kitex_gen/user"
	"testing"
)

func TestUserService_Register(t *testing.T) {
	s := NewUserService(context.Background())

	request := &user.UserRegisterReq{}
	resp, err := s.Register(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// TEST: edit your unit test

}
