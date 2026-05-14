package boards_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *BoardRepositoryImpl) SaveBoard(ctx context.Context, board domain.Board) (domain.Board, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE boards
	SET
		title = $2,
		space_id = $3
	WHERE id = $1
	RETURNING id, title, space_id, created_at;
	`

	row := r.pool.QueryRow(ctx, query, board.ID, board.Title, board.SpaceID)

	var boardModel BoardModel
	err := row.Scan(
		&boardModel.ID,
		&boardModel.Title,
		&boardModel.SpaceID,
		&boardModel.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Board{}, fmt.Errorf("board with id='%s': %w", board.ID, core_errors.ErrNotFound)
		}
		return domain.Board{}, fmt.Errorf("scan error: %w", err)
	}

	boardDomain := domainFromModel(boardModel)
	return boardDomain, nil
}
