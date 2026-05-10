package spaces_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *SpaceRepositoryImpl) SaveSpace(ctx context.Context, space domain.Space) (domain.Space, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE spaces
	SET
		title = $2,
		description = $3,
		owner_id = $4,
		is_archived = $5
	WHERE id = $1
	RETURNING id, title, description, owner_id, created_at, is_archived
	`

	row := r.pool.QueryRow(ctx, query, space.ID, space.Title, space.Description, space.OwnerID, space.IsArchived)

	var spaceModel SpaceModel
	err := row.Scan(
		&spaceModel.ID,
		&spaceModel.Title,
		&spaceModel.Description,
		&spaceModel.OwnerID,
		&spaceModel.CreatedAt,
		&spaceModel.IsArchived,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Space{}, fmt.Errorf("space with id %s: %w", space.ID.String(), core_errors.ErrNotFound)
		}
		return domain.Space{}, fmt.Errorf("scan error: %w", err)
	}

	spaceDomain := domainFromModel(spaceModel)
	return spaceDomain, nil
}
