package service

import (
	"context"
)

type AuthService struct {
	ctx context.Context
}

// NewAuthService new AuthService
func NewAuthService(ctx context.Context) *AuthService {
	return &AuthService{ctx: ctx}
}
