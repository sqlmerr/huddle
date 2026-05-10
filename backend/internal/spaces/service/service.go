package spaces_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	spaces_postgres_repository "github.com/sqlmerr/huddle/backend/internal/spaces/repository/postgres"
)

type SpaceService interface {
	CreateSpace(ctx context.Context, input CreateSpaceInput) (domain.Space, error)
	GetSpace(ctx context.Context, userID uuid.UUID, spaceID uuid.UUID) (domain.Space, error)
	GetUserSpaces(ctx context.Context, userID uuid.UUID) ([]domain.Space, error)
	PatchSpace(ctx context.Context, input PatchSpaceInput) (domain.Space, error)

	// ArchiveSpace(ctx context.Context, userID uuid.UUID, spaceID uuid.UUID) error
}

type SpaceServiceImpl struct {
	repo spaces_postgres_repository.SpaceRepository
}

func NewSpaceService(repo spaces_postgres_repository.SpaceRepository) *SpaceServiceImpl {
	return &SpaceServiceImpl{repo}
}
