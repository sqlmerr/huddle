package lists_postgres_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

type ListRepository interface {
	CreateList(ctx context.Context, list domain.List) (domain.List, error)
	GetList(ctx context.Context, listID uuid.UUID) (domain.List, error)
	GetBoardLists(ctx context.Context, boardID uuid.UUID) ([]domain.List, error)
	SaveList(ctx context.Context, list domain.List) (domain.List, error)
	DeleteList(ctx context.Context, listID uuid.UUID) error
}

type ListRepositoryImpl struct {
	pool core_postgres_pool.Pool
}

func NewListRepository(pool core_postgres_pool.Pool) *ListRepositoryImpl {
	return &ListRepositoryImpl{pool}
}
