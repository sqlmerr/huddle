package boards_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *BoardServiceImpl) CreateBoard(ctx context.Context, userID uuid.UUID, input CreateBoardInput) (domain.Board, error) {
	if err := s.accessService.CanAccessSpaceByID(ctx, userID, input.SpaceID); err != nil {
		return domain.Board{}, fmt.Errorf("access denied: %w", err)
	}

	boardDomain := domain.NewBoardUninitialized(input.Title, input.SpaceID)
	board, err := s.boardRepository.CreateBoard(ctx, boardDomain)
	if err != nil {
		return domain.Board{}, fmt.Errorf("create board: %w", err)
	}

	return board, nil
}
