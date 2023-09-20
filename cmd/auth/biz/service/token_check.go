package service

import (
	"github.com/Ocyss/Qblog/kitex_gen/auth"
	"github.com/Ocyss/Qblog/kitex_gen/common"
	"github.com/Ocyss/Qblog/pkg/utils/tokens"
)

func (s *AuthService) TokenCheck(request *auth.AuthTokenCheckReq) (resp *common.EmptyStruct, err error) {
	resp = new(common.EmptyStruct)
	_, err = tokens.CheckToken(request.Token)
	return
}
