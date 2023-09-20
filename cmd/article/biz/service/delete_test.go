package service

import (
	"context"
	"testing"

	article "github.com/Ocyss/Qblog/kitex_gen/article"
)

func TestArticleService_Delete(t *testing.T) {
	s := NewArticleService(context.Background())

	request := &article.ArticleDeleteReq{}
	resp, err := s.Delete(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test
}
