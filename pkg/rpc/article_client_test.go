package rpc

import (
	"context"
	"testing"

	"github.com/Ocyss/Qblog/kitex_gen/user"

	"github.com/Ocyss/Qblog/kitex_gen/article"
)

func TestArtcleClient(t *testing.T) {
	var resp any
	var err error
	client := NewArticleClient()
	userClient := NewUserClient()
	ctx := context.Background()
	resp, err = client.Get(ctx, &article.ArticleReq{Aid: 645})
	Logs(t, resp, err)
	resp, err = client.GetList(ctx, &article.ArticlesReq{})
	Logs(t, resp, err)
	userResp, err := userClient.Register(ctx, &user.UserRegisterReq{Username: "ArtcleTest1", Password: "testtest"})
	defer func() {
		resp, err = userClient.Delete(ctx, &user.UserDeleteReq{Username: "ArtcleTest1", Password: "testtest"})
		Log(t, nil, err)
	}()
	Log(t, nil, err)
	ArticleResp, err := client.Create(ctx, &article.ArticleCreateReq{
		Userid: userResp.Uid,
		Data: &article.ArticleEdit{
			Title:     "一个test用例罢了",
			Introduce: "简介：测试",
			Content:   "正文",
			Tags:      []string{"test", "test1", "test2"},
			Show:      true,
		},
	})
	defer func() {
		del := true
		resp, err = client.Delete(ctx, &article.ArticleDeleteReq{Aid: ArticleResp.Aid, IsDelete: &del})
		Log(t, nil, err)
	}()
	Log(t, ArticleResp, err)
	resp, err = client.Edit(ctx, &article.ArticleEditReq{Uri: ArticleResp.Uri, Data: &article.ArticleEdit{
		Title:     "修改标题咯",
		Introduce: "",
		Content:   "测试正文",
		Tags:      []string{"测试", "test", "test2"},
		Show:      false,
	}})
	Log(t, nil, err)
	resp, err = client.Get(ctx, &article.ArticleReq{Aid: ArticleResp.Aid})
	Logs(t, resp, err)
	resp, err = client.GetList(ctx, &article.ArticlesReq{})
	Logs(t, resp, err)
}
