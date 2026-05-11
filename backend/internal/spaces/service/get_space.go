package spaces_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *SpaceServiceImpl) GetSpace(ctx context.Context, userID uuid.UUID, spaceID uuid.UUID) (domain.Space, error) {
	space, err := s.repo.GetSpace(ctx, spaceID)
	if err != nil {
		return domain.Space{}, fmt.Errorf("get space: %w", err)
	}

	if err := s.accessService.CanAccessSpace(ctx, userID, space); err != nil {
		return domain.Space{}, fmt.Errorf("unable to access the space with id='%s': %w", spaceID.String(), err)
	}

	return space, nil
}
