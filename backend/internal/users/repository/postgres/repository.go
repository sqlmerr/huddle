package users_postgres_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

type UserRepositoryImpl struct {
	pool core_postgres_pool.Pool
}

type UserRepository interface {
	GetUser(ctx context.Context, userID uuid.UUID) (domain.User, error)
}

func NewUserRepository(pool core_postgres_pool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{pool: pool}
}
