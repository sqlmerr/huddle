package users_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type UserService struct {
	repo UserRepository
}

type UserRepository interface {
	GetUser(ctx context.Context, userID uuid.UUID) (domain.User, error)
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}
