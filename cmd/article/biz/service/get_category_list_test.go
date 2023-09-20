package service

import (
	"context"
	"testing"
)

func TestArticleService_GetCategoryList(t *testing.T) {
	s := NewArticleService(context.Background())
	resp, err := s.GetCategoryList()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test
}
