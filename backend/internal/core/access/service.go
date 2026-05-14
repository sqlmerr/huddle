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
	listRepository  ListRepository
}

func NewAccessService(spaceRepo SpaceRepository, boardRepo BoardRepository, listRepo ListRepository) *AccessServiceImpl {
	return &AccessServiceImpl{
		spaceRepository: spaceRepo,
		boardRepository: boardRepo,
		listRepository:  listRepo,
	}
}

func (s *AccessServiceImpl) CanAccessSpace(ctx context.Context, userID uuid.UUID, space domain.Space) error {
	// TODO: permissions and space members
	if space.OwnerID != userID {
		return fmt.Errorf(
			"space with id='%s': %w",
			space.ID,
			core_errors.ErrAccessDenied,
		)
	}

	return nil
}

func (s *AccessServiceImpl) CanAccessSpaceByID(ctx context.Context, userID, spaceID uuid.UUID) error {
	space, err := s.spaceRepository.GetSpace(ctx, spaceID)
	if err != nil {
		return fmt.Errorf(
			"space with id='%s': %w",
			spaceID,
			err,
		)
	}

	return s.CanAccessSpace(ctx, userID, space)
}

func (s *AccessServiceImpl) CanAccessBoardByID(ctx context.Context, userID, boardID uuid.UUID) error {
	board, err := s.boardRepository.GetBoard(ctx, boardID)
	if err != nil {
		return fmt.Errorf(
			"board with id='%s': %w",
			boardID,
			err,
		)
	}

	return s.CanAccessSpaceByID(ctx, userID, board.SpaceID)
}

func (s *AccessServiceImpl) CanAccessBoard(ctx context.Context, userID uuid.UUID, board domain.Board) error {
	return s.CanAccessSpaceByID(ctx, userID, board.SpaceID)
}

func (s *AccessServiceImpl) CanAccessList(ctx context.Context, userID uuid.UUID, list domain.List) error {
	return s.CanAccessBoardByID(ctx, userID, list.BoardID)
}

func (s *AccessServiceImpl) CanAccessListByID(ctx context.Context, userID, listID uuid.UUID) error {
	list, err := s.listRepository.GetList(ctx, listID)
	if err != nil {
		return fmt.Errorf(
			"list with id='%s': %w",
			listID,
			err,
		)
	}
	return s.CanAccessList(ctx, userID, list)
}
