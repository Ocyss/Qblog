package service

import (
	common "github.com/Ocyss/Qblog/kitex_gen/common"
	user "github.com/Ocyss/Qblog/kitex_gen/user"
)

func (s *UserService) Logout(request *user.UserLogoutReq) (resp *common.EmptyStruct, err error) {
	resp = new(common.EmptyStruct)
	// TODO UserService.Logout Finish your business logic.

	return
}
