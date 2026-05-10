package spaces_postgres_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

type SpaceRepository interface {
	CreateSpace(ctx context.Context, space domain.Space) (domain.Space, error)
	GetSpace(ctx context.Context, spaceID uuid.UUID) (domain.Space, error)
	GetSpacesByOwnerID(ctx context.Context, ownerID uuid.UUID, filter GetSpacesByOwnerIDFilter) ([]domain.Space, error)
	SaveSpace(ctx context.Context, space domain.Space) (domain.Space, error)

	// TODO: count spaces by owner id
}

type SpaceRepositoryImpl struct {
	pool core_postgres_pool.Pool
}

func NewSpaceRepository(pool core_postgres_pool.Pool) *SpaceRepositoryImpl {
	return &SpaceRepositoryImpl{pool}
}
