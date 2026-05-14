package lists_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *ListRepositoryImpl) CreateList(ctx context.Context, list domain.List) (domain.List, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO lists (title, board_id, position) VALUES ($1, $2, $3)
	RETURNING id, title, board_id, position, created_at;
	`

	row := r.pool.QueryRow(ctx, query, list.Title, list.BoardID, list.Position)

	var listModel ListModel
	err := row.Scan(
		&listModel.ID,
		&listModel.Title,
		&listModel.BoardID,
		&listModel.Position,
		&listModel.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.List{}, fmt.Errorf(
				"board with id='%s': %w",
				list.BoardID,
				core_errors.ErrNotFound,
			)
		}
		return domain.List{}, fmt.Errorf("scan error: %w", err)
	}

	listDomain := domainFromModel(listModel)
	return listDomain, nil
}
