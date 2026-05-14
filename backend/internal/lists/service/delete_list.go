package lists_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *ListServiceImpl) DeleteList(ctx context.Context, userID uuid.UUID, listID uuid.UUID) error {
	if err := s.accessService.CanAccessListByID(ctx, userID, listID); err != nil {
		return fmt.Errorf("cannot access list: %w", err)
	}

	err := s.repo.DeleteList(ctx, listID)
	if err != nil {
		return fmt.Errorf("delete list: %w", err)
	}
	return nil
}
