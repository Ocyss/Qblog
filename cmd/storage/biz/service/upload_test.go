package service

import (
	"context"
	"testing"

	"github.com/Ocyss/Qblog/kitex_gen/storage"
)

func TestStorageService_Upload(t *testing.T) {
	s := NewStorageService(context.Background())

	request := &storage.StorageUploadReq{}
	resp, err := s.Upload(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test
}
