package service

import (
	"context"
	search "github.com/Ocyss/Qblog/kitex_gen/search"
	"testing"
)

func TestSearchService_FindArticle(t *testing.T) {
	s := NewSearchService(context.Background())

	request := &search.SearchArticleReq{}
	resp, err := s.FindArticle(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// TEST: edit your unit test

}
