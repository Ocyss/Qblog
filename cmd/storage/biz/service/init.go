package service

import (
	"context"
)

type StorageService struct {
	ctx context.Context
}

// NewStorageService new StorageService
func NewStorageService(ctx context.Context) *StorageService {
	return &StorageService{ctx: ctx}
}
