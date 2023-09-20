package main

import (
	"context"

	"github.com/Ocyss/Qblog/kitex_gen/storage"

	"github.com/Ocyss/Qblog/cmd/storage/biz/service"
)

// StorageServiceImpl implements the last service interface defined in the IDL.
type StorageServiceImpl struct{}

// Upload implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) Upload(ctx context.Context, request *storage.StorageUploadReq) (resp *storage.StorageUploadResp, err error) {
	storageService := service.NewStorageService(ctx)
	resp, err = storageService.Upload(request)
	return
}
