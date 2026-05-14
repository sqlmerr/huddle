package lists_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (r *ListRepositoryImpl) GetBoardLists(ctx context.Context, boardID uuid.UUID) ([]domain.List, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, title, board_id, position, created_at
	FROM lists
	WHERE board_id = $1
	ORDER BY position;
	`

	rows, err := r.pool.Query(ctx, query, boardID)
	if err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}
	var listModels []ListModel
	for rows.Next() {
		var listModel ListModel
		err := rows.Scan(
			&listModel.ID,
			&listModel.Title,
			&listModel.BoardID,
			&listModel.Position,
			&listModel.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		listModels = append(listModels, listModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	listDomains := domainFromModels(listModels)
	return listDomains, nil
}
