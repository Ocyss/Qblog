package rpc

import (
	"context"
	"testing"

	"github.com/Ocyss/Qblog/kitex_gen/user"
)

func TestUserClient(t *testing.T) {
	var resp any
	var err error
	client := NewUserClient()
	ctx := context.Background()
	resp, err = client.Login(ctx, &user.UserLoginReq{Username: "adminTest", Password: "123"})
	Logs(t, resp, err)
	resp, err = client.Register(ctx, &user.UserRegisterReq{Username: "adminTest", Nickname: "666", Password: "123"})
	defer func() {
		resp, err = client.Delete(ctx, &user.UserDeleteReq{Username: "adminTest", Password: "123"})
		Log(t, resp, err)
	}()
	Log(t, resp, err)
	resp, err = client.Login(ctx, &user.UserLoginReq{Username: "adminTest", Password: "123"})
	Log(t, resp, err)
}
