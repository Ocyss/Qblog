package service

import (
	"context"
	"testing"

	comment "github.com/Ocyss/Qblog/kitex_gen/comment"
)

func TestCommentService_Create(t *testing.T) {
	s := NewCommentService(context.Background())

	request := &comment.CommentCreateReq{}
	resp, err := s.Create(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// TEST: edit your unit test
}
