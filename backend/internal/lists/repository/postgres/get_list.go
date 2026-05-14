package lists_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *ListRepositoryImpl) GetList(ctx context.Context, listID uuid.UUID) (domain.List, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, title, board_id, position, created_at
	FROM lists
	WHERE id = $1;
	`

	row := r.pool.QueryRow(ctx, query, listID)

	var listModel ListModel
	err := row.Scan(
		&listModel.ID,
		&listModel.Title,
		&listModel.BoardID,
		&listModel.Position,
		&listModel.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.List{}, fmt.Errorf("list with id='%s': %w", listID, core_errors.ErrNotFound)
		}
		return domain.List{}, fmt.Errorf("scan error: %w", err)
	}

	listDomain := domainFromModel(listModel)
	return listDomain, nil
}
