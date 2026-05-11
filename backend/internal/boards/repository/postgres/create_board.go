package boards_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *BoardRepositoryImpl) CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO boards (title, space_id) VALUES ($1, $2)
	RETURNING id, title, space_id, created_at;
	`

	row := r.pool.QueryRow(ctx, query, board.Title, board.SpaceID)

	var boardModel BoardModel
	err := row.Scan(
		&boardModel.ID,
		&boardModel.Title,
		&boardModel.SpaceID,
		&boardModel.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Board{}, fmt.Errorf(
				"space with id='%s': %w",
				board.SpaceID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Board{}, fmt.Errorf("scan: %w", err)
	}

	boardDomain := domainFromModel(boardModel)
	return boardDomain, nil
}
