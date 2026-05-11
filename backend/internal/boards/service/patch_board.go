package boards_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *BoardServiceImpl) PatchBoard(ctx context.Context, userID uuid.UUID, input PatchBoardInput) (domain.Board, error) {
	board, err := s.boardRepository.GetBoard(ctx, input.BoardID)
	if err != nil {
		return domain.Board{}, fmt.Errorf("get board: %w", err)
	}

	if err := s.accessService.CanAccessBoard(ctx, userID, board); err != nil {
		return domain.Board{}, fmt.Errorf("access denied: %w", err)
	}

	patchedBoard, err := input.ApplyPatch(board)
	if err != nil {
		return domain.Board{}, fmt.Errorf("apply patch: %w", err)
	}

	boardDomain, err := s.boardRepository.SaveBoard(ctx, patchedBoard)
	if err != nil {
		return domain.Board{}, fmt.Errorf("save board: %w", err)
	}

	return boardDomain, nil
}
