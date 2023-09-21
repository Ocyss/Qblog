package service

import (
	"context"
	comment "github.com/Ocyss/Qblog/kitex_gen/comment"
	"testing"
)

func TestCommentService_Get(t *testing.T) {
	s := NewCommentService(context.Background())

	request := &comment.CommentListReq{}
	resp, err := s.Get(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// TEST: edit your unit test

}
