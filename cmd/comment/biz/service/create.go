package service

import (
	"github.com/Ocyss/Qblog/kitex_gen/comment"
	"github.com/Ocyss/Qblog/kitex_gen/common"
)

func (s *CommentService) Create(request *comment.CommentCreateReq) (resp *common.EmptyStruct, err error) {
	resp = new(common.EmptyStruct)
	// TODO CommentService.Create Finish your business logic.

	return
}
