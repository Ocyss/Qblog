package service

import (
	user "github.com/Ocyss/Qblog/kitex_gen/user"
)

func (s *UserService) Register(request *user.UserRegisterReq) (resp *user.UserRegisterResp, err error) {
	resp = new(user.UserRegisterResp)
	// TODO UserService.Register Finish your business logic.

	return
}
