package service

import (
	"context"

	"vananin/go-from-scratch/internal/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	// Business rules can be added here later.

	return s.repo.GetByID(ctx, id)
}
