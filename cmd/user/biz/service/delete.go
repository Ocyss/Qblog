package service

import (
	"github.com/Ocyss/Qblog/kitex_gen/common"
	"github.com/Ocyss/Qblog/kitex_gen/user"
)

func (s *UserService) Delete(request *user.UserDeleteReq) (resp *common.EmptyStruct, err error) {
	resp = new(common.EmptyStruct)
	// TODO UserService.Delete Finish your business logic.

	return
}
