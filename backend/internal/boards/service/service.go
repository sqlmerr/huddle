package boards_service

import (
	"context"

	"github.com/google/uuid"
	boards_postgres_repository "github.com/sqlmerr/huddle/backend/internal/boards/repository/postgres"
	core_access "github.com/sqlmerr/huddle/backend/internal/core/access"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type BoardService interface {
	CreateBoard(ctx context.Context, userID uuid.UUID, input CreateBoardInput) (domain.Board, error)
	GetBoard(ctx context.Context, userID uuid.UUID, boardID uuid.UUID) (domain.Board, error)
	GetSpaceBoards(ctx context.Context, userID uuid.UUID, spaceID uuid.UUID) ([]domain.Board, error)
	PatchBoard(ctx context.Context, userID uuid.UUID, input PatchBoardInput) (domain.Board, error)
	DeleteBoard(ctx context.Context, userID uuid.UUID, boardID uuid.UUID) error
}

type BoardServiceImpl struct {
	boardRepository boards_postgres_repository.BoardRepository
	accessService   core_access.AccessService
}

func NewBoardService(
	boardRepo boards_postgres_repository.BoardRepository,
	accessService core_access.AccessService,
) *BoardServiceImpl {
	return &BoardServiceImpl{boardRepo, accessService}
}
