package boards_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *BoardServiceImpl) DeleteBoard(ctx context.Context, userID uuid.UUID, boardID uuid.UUID) error {
	if err := s.accessService.CanAccessBoardByID(ctx, userID, boardID); err != nil {
		return fmt.Errorf("access denied: %w", err)
	}

	err := s.boardRepository.DeleteBoard(ctx, boardID)
	if err != nil {
		return fmt.Errorf("delete board: %w", err)
	}
	return nil
}
