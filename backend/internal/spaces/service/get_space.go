package spaces_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (s *SpaceServiceImpl) GetSpace(ctx context.Context, userID uuid.UUID, spaceID uuid.UUID) (domain.Space, error) {
	space, err := s.repo.GetSpace(ctx, spaceID)
	if err != nil {
		return domain.Space{}, fmt.Errorf("get space: %w", err)
	}

	if space.OwnerID != userID { // TODO: space members
		return domain.Space{}, fmt.Errorf("unable to access the space with id='%s': %w", spaceID.String(), core_errors.ErrAccessDenied)
	}

	return space, nil
}
