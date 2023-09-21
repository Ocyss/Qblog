package rpc

import "github.com/Ocyss/Qblog/pkg/rpc"

var Article *rpc.ArticleClient

func Init() {
	Article = rpc.NewArticleClient()
}
