package main

import (
	"context"

	"github.com/Ocyss/Qblog/kitex_gen/comment"
	"github.com/Ocyss/Qblog/kitex_gen/common"

	"github.com/Ocyss/Qblog/cmd/comment/biz/service"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// Create implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Create(ctx context.Context, request *comment.CommentCreateReq) (resp *common.EmptyStruct, err error) {
	commentService := service.NewCommentService(ctx)
	resp, err = commentService.Create(request)
	return
}
