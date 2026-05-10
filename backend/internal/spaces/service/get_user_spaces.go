package spaces_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	spaces_postgres_repository "github.com/sqlmerr/huddle/backend/internal/spaces/repository/postgres"
)

func (s *SpaceServiceImpl) GetUserSpaces(ctx context.Context, userID uuid.UUID) ([]domain.Space, error) {
	return s.getUserSpaces(ctx, userID, false)
}

func (s *SpaceServiceImpl) GetArchivedUserSpaces(ctx context.Context, userID uuid.UUID) ([]domain.Space, error) {
	return s.getUserSpaces(ctx, userID, true)
}

func (s *SpaceServiceImpl) getUserSpaces(ctx context.Context, userID uuid.UUID, isArchived bool) ([]domain.Space, error) {
	spaces, err := s.repo.GetSpacesByOwnerID(ctx, userID, spaces_postgres_repository.GetSpacesByOwnerIDFilter{IsArchived: &isArchived})
	if err != nil {
		return nil, fmt.Errorf("get spaces by user id='%s': %w", userID.String(), err)
	}

	return spaces, nil
}
