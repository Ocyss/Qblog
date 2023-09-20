package rpc

import (
	"time"

	"github.com/Ocyss/Qblog/pkg/configs"
	"github.com/Ocyss/Qblog/pkg/middleware"

	"github.com/Ocyss/Qblog/kitex_gen/article/articleservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type ArticleClient struct {
	articleservice.Client
}

func NewArticleClient() *ArticleClient {
	r, err := etcd.NewEtcdResolver(configs.GetRegistry().Address)
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err.Error())
	}

	c, err := articleservice.NewClient("article",
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMiddleware(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(5*time.Second),
		client.WithConnectTimeout(100*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		klog.Fatalf("new user client failed: %v", err.Error())
	}

	return &ArticleClient{Client: c}
}
