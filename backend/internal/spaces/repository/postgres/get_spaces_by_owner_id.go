package spaces_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type GetSpacesByOwnerIDFilter struct {
	IsArchived *bool
}

func (r *SpaceRepositoryImpl) GetSpacesByOwnerID(ctx context.Context, ownerID uuid.UUID, filters GetSpacesByOwnerIDFilter) ([]domain.Space, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, title, description, owner_id, created_at, is_archived
	FROM spaces
	WHERE owner_id = $1 %s
	ORDER BY created_at;
	`

	args := []any{ownerID}
	if filters.IsArchived != nil {
		query = fmt.Sprintf(query, "AND is_archived = $2")
		args = append(args, *filters.IsArchived)
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("get spaces: %w", err)
	}
	defer rows.Close()

	var spaceModels []SpaceModel

	for rows.Next() {
		var spaceModel SpaceModel

		err := rows.Scan(
			&spaceModel.ID,
			&spaceModel.Title,
			&spaceModel.Description,
			&spaceModel.OwnerID,
			&spaceModel.CreatedAt,
			&spaceModel.IsArchived,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		spaceModels = append(spaceModels, spaceModel)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	domains := domainFromModels(spaceModels)
	return domains, nil
}
