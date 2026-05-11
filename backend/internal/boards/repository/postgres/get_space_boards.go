package boards_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (r *BoardRepositoryImpl) GetSpaceBoards(ctx context.Context, spaceID uuid.UUID) ([]domain.Board, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, title, space_id, created_at
	FROM boards
	WHERE space_id = $1
	ORDER BY created_at;
	`

	rows, err := r.pool.Query(ctx, query, spaceID)
	if err != nil {
		return nil, fmt.Errorf("get boards: %w", err)
	}

	var boardsModels []BoardModel
	for rows.Next() {
		var boardModel BoardModel
		err := rows.Scan(
			&boardModel.ID,
			&boardModel.Title,
			&boardModel.SpaceID,
			&boardModel.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		boardsModels = append(boardsModels, boardModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	domains := domainFromModels(boardsModels)
	return domains, nil
}
