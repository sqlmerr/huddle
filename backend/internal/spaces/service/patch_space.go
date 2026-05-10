package spaces_service

import (
	"context"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (s *SpaceServiceImpl) PatchSpace(ctx context.Context, input PatchSpaceInput) (domain.Space, error) {
	space, err := s.repo.GetSpace(ctx, input.SpaceID)
	if err != nil {
		return domain.Space{}, fmt.Errorf("get space: %w", err)
	}

	if space.OwnerID != input.UserID { // TODO: space members
		return domain.Space{}, fmt.Errorf("unable to access the space with id='%s': %w", input.SpaceID.String(), core_errors.ErrAccessDenied)
	}

	patchedSpace, err := input.ApplyPatch(space)
	if err != nil {
		return domain.Space{}, fmt.Errorf("apply patch: %w", err)
	}

	spaceDomain, err := s.repo.SaveSpace(ctx, patchedSpace)
	if err != nil {
		return domain.Space{}, fmt.Errorf("save space: %w", err)
	}
	return spaceDomain, nil
}
