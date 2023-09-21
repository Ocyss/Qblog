package service

import (
	"context"
	article "github.com/Ocyss/Qblog/kitex_gen/article"
	"testing"
)

func TestArticleService_CheckIdUri(t *testing.T) {
	s := NewArticleService(context.Background())

	request := &article.ArticleIdUriReq{}

	err := s.CheckIdUri(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	// TEST: edit your unit test

}
