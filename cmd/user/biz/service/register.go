package service

import (
	"github.com/Ocyss/Qblog/cmd/user/biz/dal/db"
	"github.com/Ocyss/Qblog/kitex_gen/user"
	"github.com/Ocyss/Qblog/pkg/constants"
	"github.com/Ocyss/Qblog/pkg/utils/tokens"
	"gopkg.in/hlandau/passlib.v1"
)

func (s *UserService) Register(request *user.UserRegisterReq) (resp *user.UserRegisterResp, err error) {
	newPswd, err := passlib.Hash(request.Password)
	if err != nil {
		return
	}
	uid, err := db.CreateUser(s.ctx, &db.User{
		Username: request.Username,
		Password: newPswd,
		Nickname: request.Nickname,
		Role:     constants.Guest,
	})
	if err != nil {
		return
	}
	token, err := tokens.GetToken(uid, request.Username)
	if err != nil {
		return
	}
	resp = &user.UserRegisterResp{
		Token: token,
		Uid:   int32(uid),
	}
	return
}
