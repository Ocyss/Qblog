package main

import (
	"context"

	"github.com/Ocyss/Qblog/cmd/auth/biz/service"
	"github.com/Ocyss/Qblog/kitex_gen/auth"
	"github.com/Ocyss/Qblog/kitex_gen/common"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// TokenCheck implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) TokenCheck(ctx context.Context, request *auth.AuthTokenCheckReq) (resp *common.EmptyStruct, err error) {
	authService := service.NewAuthService(ctx)
	resp, err = authService.TokenCheck(request)
	return
}
