package core_access

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

type AccessServiceImpl struct {
	spaceRepository SpaceRepository
	boardRepository BoardRepository
}

func NewAccessService(spaceRepo SpaceRepository, boardRepo BoardRepository) *AccessServiceImpl {
	return &AccessServiceImpl{
		spaceRepository: spaceRepo,
		boardRepository: boardRepo,
	}
}

func (s *AccessServiceImpl) CanAccessSpace(ctx context.Context, userID, spaceID uuid.UUID) error {
	space, err := s.spaceRepository.GetSpace(ctx, spaceID)
	if err != nil {
		return fmt.Errorf(
			"space with id='%s': %w",
			spaceID,
			core_errors.ErrNotFound,
		)
	}
	// TODO: permissions and space members
	if space.OwnerID != userID {
		return fmt.Errorf(
			"space with id='%s': %w",
			spaceID,
			core_errors.ErrAccessDenied,
		)
	}

	return nil
}

func (s *AccessServiceImpl) CanAccessBoardByID(ctx context.Context, userID, boardID uuid.UUID) error {
	board, err := s.boardRepository.GetBoard(ctx, boardID)
	if err != nil {
		return fmt.Errorf(
			"board with id='%s': %w",
			boardID,
			core_errors.ErrNotFound,
		)
	}

	return s.CanAccessSpace(ctx, userID, board.SpaceID)
}

func (s *AccessServiceImpl) CanAccessBoard(ctx context.Context, userID uuid.UUID, board domain.Board) error {
	return s.CanAccessSpace(ctx, userID, board.SpaceID)
}
