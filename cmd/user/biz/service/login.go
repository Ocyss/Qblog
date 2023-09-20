package service

import (
	user "github.com/Ocyss/Qblog/kitex_gen/user"
)

func (s *UserService) Login(request *user.UserLoginReq) (resp *user.UserLoginResp, err error) {
	resp = new(user.UserLoginResp)
	// TODO UserService.Login Finish your business logic.

	return
}
