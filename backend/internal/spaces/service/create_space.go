package spaces_service

import (
	"context"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *SpaceServiceImpl) CreateSpace(ctx context.Context, input CreateSpaceInput) (domain.Space, error) {
	spaceDomain := domain.NewSpaceUninitialized(input.Title, input.Description, input.OwnerID)
	if err := spaceDomain.Validate(); err != nil {
		return domain.Space{}, fmt.Errorf("validate input: %w", err)
	}

	// TODO: check the number of spaces created by the user

	space, err := s.repo.CreateSpace(ctx, spaceDomain)
	if err != nil {
		return domain.Space{}, fmt.Errorf("create space: %w", err)
	}

	return space, nil
}
