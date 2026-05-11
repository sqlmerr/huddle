package boards_postgres_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

type BoardRepository interface {
	CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error)
	GetBoard(ctx context.Context, boardID uuid.UUID) (domain.Board, error)
	GetSpaceBoards(ctx context.Context, spaceID uuid.UUID) ([]domain.Board, error)
	SaveBoard(ctx context.Context, board domain.Board) (domain.Board, error)
	DeleteBoard(ctx context.Context, boardID uuid.UUID) error
}

type BoardRepositoryImpl struct {
	pool core_postgres_pool.Pool
}

func NewBoardRepository(pool core_postgres_pool.Pool) *BoardRepositoryImpl {
	return &BoardRepositoryImpl{pool}
}
