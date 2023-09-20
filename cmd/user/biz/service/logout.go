package service

import (
	"github.com/Ocyss/Qblog/cmd/user/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/common"
	"github.com/Ocyss/Qblog/kitex_gen/user"
	"github.com/Ocyss/Qblog/pkg/utils/tokens"
)

func (s *UserService) Logout(request *user.UserLogoutReq) (resp *common.EmptyStruct, err error) {
	claims, err := tokens.CheckToken(request.Token)
	if err != nil {
		return
	}
	err = db.DeleteUser(s.ctx, claims.Username, false)
	return
}
