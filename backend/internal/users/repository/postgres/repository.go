package users_postgres_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

type UsersRepositoryImpl struct {
	pool core_postgres_pool.Pool
}

type UserRepository interface {
	GetUser(ctx context.Context, userID uuid.UUID) (domain.User, error)
}

func NewUsersRepository(pool core_postgres_pool.Pool) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{pool: pool}
}
