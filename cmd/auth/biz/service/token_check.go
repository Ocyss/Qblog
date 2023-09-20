package service

import (
	"github.com/Ocyss/Qblog/kitex_gen/auth"
	"github.com/Ocyss/Qblog/kitex_gen/common"
)

func (s *AuthService) TokenCheck(request *auth.AuthTokenCheckReq) (resp *common.EmptyStruct, err error) {
	resp = new(common.EmptyStruct)
	// TODO AuthService.TokenCheck Finish your business logic.

	return
}
