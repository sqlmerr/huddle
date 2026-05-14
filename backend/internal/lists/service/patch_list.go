package lists_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *ListServiceImpl) PatchList(ctx context.Context, userID uuid.UUID, input PatchListInput) (domain.List, error) {
	list, err := s.repo.GetList(ctx, input.ListID)
	if err != nil {
		return domain.List{}, fmt.Errorf("get list: %w", err)
	}

	if err := s.accessService.CanAccessList(ctx, userID, list); err != nil {
		return domain.List{}, fmt.Errorf("cannot access list: %w", err)
	}

	patchedList, err := input.ApplyPatch(list)
	if err != nil {
		return domain.List{}, fmt.Errorf("failed to apply patch: %w", err)
	}

	listDomain, err := s.repo.SaveList(ctx, patchedList)
	if err != nil {
		return domain.List{}, fmt.Errorf("save list: %w", err)
	}

	return listDomain, nil
}
