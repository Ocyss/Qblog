package service

import (
	"context"
	article "github.com/Ocyss/Qblog/kitex_gen/article"
	"testing"
)

func TestArticleService_Get(t *testing.T) {
	s := NewArticleService(context.Background())

	request := &article.ArticleReq{}
	resp, err := s.Get(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
