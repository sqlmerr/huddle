package boards_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *BoardServiceImpl) GetBoard(ctx context.Context, userID uuid.UUID, boardID uuid.UUID) (domain.Board, error) {
	board, err := s.boardRepository.GetBoard(ctx, boardID)
	if err != nil {
		return domain.Board{}, fmt.Errorf("get board: %w", err)
	}

	if err := s.accessService.CanAccessBoard(ctx, userID, board); err != nil {
		return domain.Board{}, fmt.Errorf("access denied: %w", err)
	}

	return board, nil
}
