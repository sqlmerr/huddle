package boards_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *BoardServiceImpl) GetSpaceBoards(ctx context.Context, userID uuid.UUID, spaceID uuid.UUID) ([]domain.Board, error) {
	if err := s.accessService.CanAccessSpaceByID(ctx, userID, spaceID); err != nil {
		return nil, fmt.Errorf("access denied: %w", err)
	}

	boards, err := s.boardRepository.GetSpaceBoards(ctx, spaceID)
	if err != nil {
		return nil, fmt.Errorf("get space boards: %w", err)
	}

	return boards, nil
}
