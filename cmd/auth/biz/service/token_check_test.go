package service

import (
	"context"
	"testing"

	auth "github.com/Ocyss/Qblog/kitex_gen/auth"
)

func TestAuthService_TokenCheck(t *testing.T) {
	s := NewAuthService(context.Background())

	request := &auth.AuthTokenCheckReq{}
	resp, err := s.TokenCheck(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// TEST: edit your unit test
}
