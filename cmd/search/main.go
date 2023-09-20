package main

import (
	"net"

	"github.com/Ocyss/Qblog/pkg/bound"

	"github.com/Ocyss/Qblog/pkg/constants"

	tracer2 "github.com/Ocyss/Qblog/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/limit"
	trace "github.com/kitex-contrib/tracer-opentracing"

	"github.com/Ocyss/Qblog/cmd/search/biz/dal"
	"github.com/Ocyss/Qblog/cmd/search/biz/rpc"
	"github.com/Ocyss/Qblog/pkg/middleware"

	"github.com/Ocyss/Qblog/pkg/configs"

	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/Ocyss/Qblog/kitex_gen/search/searchservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func Init() {
	tracer2.InitJaeger(constants.SearchServiceName)
	dal.Init()
	rpc.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry(configs.GetRegistry().Address)
	if err != nil {
		klog.Fatalf("new registry failed: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		klog.Fatalf("resolve addr failed: %v", err)
	}

	Init()
	svr := searchservice.NewServer(new(SearchServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.SearchServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                               // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),                                             // registry
	)
	if err := svr.Run(); err != nil {
		klog.Fatalf("%s server stopped with error: %v", constants.SearchServiceName, err)
	}
}
