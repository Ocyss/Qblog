package service

import (
	"context"
	storage "github.com/Ocyss/Qblog/kitex_gen/storage"
	"testing"
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
	// TEST: edit your unit test

}
