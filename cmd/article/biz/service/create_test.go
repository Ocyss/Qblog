package service

import (
	"context"
	article "github.com/Ocyss/Qblog/kitex_gen/article"
	"testing"
)

func TestArticleService_Create(t *testing.T) {
	s := NewArticleService(context.Background())

	request := &article.ArticleCreateReq{}
	resp, err := s.Create(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
