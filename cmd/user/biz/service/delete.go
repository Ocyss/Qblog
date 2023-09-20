package service

import (
	"github.com/Ocyss/Qblog/cmd/user/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/common"
	"github.com/Ocyss/Qblog/kitex_gen/user"
	"gopkg.in/hlandau/passlib.v1"
)

func (s *UserService) Delete(request *user.UserDeleteReq) (resp *common.EmptyStruct, err error) {
	dbUser, err := db.GetUserByUserName(s.ctx, request.Username)
	if err != nil {
		return
	}
	_, err = passlib.Verify(request.Password, dbUser.Password)
	if err != nil {
		return
	}
	err = db.DeleteUser(s.ctx, request.Username, true)
	return
}
