package lists_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *ListServiceImpl) GetBoardLists(ctx context.Context, userID uuid.UUID, boardID uuid.UUID) ([]domain.List, error) {
	if err := s.accessService.CanAccessBoardByID(ctx, userID, boardID); err != nil {
		return nil, fmt.Errorf("cannot access board: %w", err)
	}

	lists, err := s.repo.GetBoardLists(ctx, boardID)
	if err != nil {
		return nil, fmt.Errorf("get board lists: %w", err)
	}
	return lists, nil
}
