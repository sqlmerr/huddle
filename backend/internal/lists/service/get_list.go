package lists_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *ListServiceImpl) GetList(ctx context.Context, userID uuid.UUID, listID uuid.UUID) (domain.List, error) {
	list, err := s.repo.GetList(ctx, listID)
	if err != nil {
		return domain.List{}, fmt.Errorf("list with id='%s': %w", listID, err)
	}

	if err := s.accessService.CanAccessList(ctx, userID, list); err != nil {
		return domain.List{}, fmt.Errorf("cannot access list: %w", err)
	}

	return list, nil
}
