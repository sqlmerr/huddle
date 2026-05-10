package spaces_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *SpaceServiceImpl) GetUserSpaces(ctx context.Context, userID uuid.UUID) ([]domain.Space, error) {
	spaces, err := s.repo.GetSpacesByOwnerID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get spaces by user id='%s': %w", userID.String(), err)
	}

	return spaces, nil
}
