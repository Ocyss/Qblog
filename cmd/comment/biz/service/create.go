package service

import (
	"github.com/Ocyss/Qblog/cmd/comment/biz/dal/db"
	"github.com/Ocyss/Qblog/cmd/comment/biz/rpc"
	"github.com/Ocyss/Qblog/kitex_gen/article"
	"github.com/Ocyss/Qblog/kitex_gen/comment"
	"github.com/Ocyss/Qblog/kitex_gen/common"
	"github.com/Ocyss/Qblog/pkg/conv"
)

func (s *CommentService) Create(request *comment.CommentCreateReq) (resp *common.EmptyStruct, err error) {
	if err = rpc.Article.CheckIdUri(s.ctx, &article.ArticleIdUriReq{Id: request.Aid}); err != nil {
		return
	}
	err = db.CreateComment(s.ctx, &db.Comment{
		CommentUser: db.CommentUser{
			Uid:   conv.ToUint(request.Data.User.Uid),
			Name:  request.Data.User.Name,
			Email: request.Data.User.Email,
			Qq:    request.Data.User.Qq,
		},
		Aid:      uint(request.Aid),
		Show:     true,
		Content:  request.Data.Content,
		ReplyID:  conv.ToUint(request.Data.Reply),
		FatherID: conv.ToUint(request.Data.Father),
	})
	return
}
