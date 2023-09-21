package rpc

import (
	"time"

	"github.com/Ocyss/Qblog/pkg/constants"

	"github.com/Ocyss/Qblog/pkg/configs"
	"github.com/Ocyss/Qblog/pkg/middleware"

	"github.com/Ocyss/Qblog/kitex_gen/comment/commentservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

type CommentClient struct {
	commentservice.Client
}

func NewCommentClient() *CommentClient {
	r, err := etcd.NewEtcdResolver(configs.GetRegistry().Address)
	if err != nil {
		klog.Fatalf("new resolver failed: %v", err.Error())
	}

	c, err := commentservice.NewClient(constants.CommentServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMiddleware(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(5*time.Second),              // rpc timeout
		client.WithConnectTimeout(100*time.Millisecond),   // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),
	)
	if err != nil {
		klog.Fatalf("new %s client failed: %v", constants.CommentServiceName, err.Error())
	}

	return &CommentClient{Client: c}
}
