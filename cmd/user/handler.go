package main

import (
	"context"
	common "github.com/Ocyss/Qblog/kitex_gen/common"
	user "github.com/Ocyss/Qblog/kitex_gen/user"

	"github.com/Ocyss/Qblog/cmd/user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.UserLoginReq) (resp *user.UserLoginResp, err error) {
	userService := service.NewUserService(ctx)
	resp, err = userService.Login(request)
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.UserRegisterReq) (resp *user.UserRegisterResp, err error) {
	userService := service.NewUserService(ctx)
	resp, err = userService.Register(request)
	return
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, request *user.UserLogoutReq) (resp *common.EmptyStruct, err error) {
	userService := service.NewUserService(ctx)
	resp, err = userService.Logout(request)
	return
}

// Delete implements the UserServiceImpl interface.
func (s *UserServiceImpl) Delete(ctx context.Context, request *user.UserDeleteReq) (resp *common.EmptyStruct, err error) {
	userService := service.NewUserService(ctx)
	resp, err = userService.Delete(request)
	return
}
