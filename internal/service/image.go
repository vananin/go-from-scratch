package service

import (
	"context"
	"vananin/go-from-scratch/internal/domain"
)

type UserStore interface {
	GetByID(ctx context.Context, id int64) (*domain.User, error)
}

type CompressService struct {
	userStore UserStore
}

func NewCompressService(userStore UserStore) *CompressService {
	return &CompressService{
		userStore: userStore,
	}
}
