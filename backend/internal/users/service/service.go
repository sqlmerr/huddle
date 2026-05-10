package users_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	users_postgres_repository "github.com/sqlmerr/huddle/backend/internal/users/repository/postgres"
)

type UserServiceImpl struct {
	repo users_postgres_repository.UserRepository
}

type UserService interface {
	GetUser(ctx context.Context, userID uuid.UUID) (domain.User, error)
}

func NewUserService(repo users_postgres_repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}
