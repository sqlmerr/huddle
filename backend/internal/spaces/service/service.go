package spaces_service

import (
	"context"

	"github.com/google/uuid"
	core_access "github.com/sqlmerr/huddle/backend/internal/core/access"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	spaces_postgres_repository "github.com/sqlmerr/huddle/backend/internal/spaces/repository/postgres"
)

type SpaceService interface {
	CreateSpace(ctx context.Context, input CreateSpaceInput) (domain.Space, error)
	GetSpace(ctx context.Context, userID uuid.UUID, spaceID uuid.UUID) (domain.Space, error)
	GetUserSpaces(ctx context.Context, userID uuid.UUID) ([]domain.Space, error)
	GetArchivedUserSpaces(ctx context.Context, userID uuid.UUID) ([]domain.Space, error)
	PatchSpace(ctx context.Context, input PatchSpaceInput) (domain.Space, error)
}

type SpaceServiceImpl struct {
	repo          spaces_postgres_repository.SpaceRepository
	accessService core_access.AccessService
}

func NewSpaceService(
	repo spaces_postgres_repository.SpaceRepository,
	accessService core_access.AccessService,
) *SpaceServiceImpl {
	return &SpaceServiceImpl{repo, accessService}
}
