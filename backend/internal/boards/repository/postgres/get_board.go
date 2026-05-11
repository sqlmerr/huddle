package boards_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *BoardRepositoryImpl) GetBoard(ctx context.Context, boardID uuid.UUID) (domain.Board, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, title, space_id, created_at
	FROM boards
	WHERE id = $1;
	`

	row := r.pool.QueryRow(ctx, query, boardID)

	var boardModel BoardModel
	err := row.Scan(
		&boardModel.ID,
		&boardModel.Title,
		&boardModel.SpaceID,
		&boardModel.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Board{}, fmt.Errorf(
				"board with id='%s': %w",
				boardID,
				core_errors.ErrNotFound,
			)
		}
		return domain.Board{}, fmt.Errorf("scan: %w", err)
	}

	boardDomain := domainFromModel(boardModel)
	return boardDomain, nil
}
