package spaces_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *SpaceRepositoryImpl) GetSpace(ctx context.Context, spaceID uuid.UUID) (domain.Space, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, title, description, owner_id, created_at, is_archived
	FROM spaces
	WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, spaceID)

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
			return domain.Space{}, fmt.Errorf("space with id='%s': %w", spaceID, core_errors.ErrNotFound)
		}
		return domain.Space{}, fmt.Errorf("scan error: %w", err)
	}

	spaceDomain := domainFromModel(spaceModel)

	return spaceDomain, nil
}
