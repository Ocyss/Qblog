package service

import (
	"context"
	"testing"

	user "github.com/Ocyss/Qblog/kitex_gen/user"
)

func TestUserService_Logout(t *testing.T) {
	s := NewUserService(context.Background())

	request := &user.UserLogoutReq{}
	resp, err := s.Logout(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test
}
