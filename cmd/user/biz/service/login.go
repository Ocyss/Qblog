package service

import (
	"github.com/Ocyss/Qblog/cmd/user/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/user"
	"github.com/Ocyss/Qblog/pkg/utils/tokens"
	"gopkg.in/hlandau/passlib.v1"
)

func (s *UserService) Login(request *user.UserLoginReq) (resp *user.UserLoginResp, err error) {
	dbUser, err := db.GetUserByUserName(s.ctx, request.Username)
	if err != nil {
		return
	}
	newHash, err := passlib.Verify(request.Password, dbUser.Password)
	if err != nil {
		return
	}
	if newHash != "" {
		err = db.UpdatePasswordHash(s.ctx, dbUser, newHash)
	}
	token, err := tokens.GetToken(dbUser.ID, dbUser.Username)
	resp = &user.UserLoginResp{
		Token: token,
		Uid:   int32(dbUser.ID),
	}
	return
}
