package spaces_postgres_repository

import (
	"context"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (r *SpaceRepositoryImpl) CreateSpace(ctx context.Context, space domain.Space) (domain.Space, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO spaces (title, description, owner_id) VALUES ($1, $2, $3)
	RETURNING id, title, description, owner_id, created_at, is_archived;
	`
	row := r.pool.QueryRow(ctx, query, space.Title, space.Description, space.OwnerID)

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
		// TODO: handle violates foreign key error
		return domain.Space{}, fmt.Errorf("scan error: %w", err)
	}

	spaceDomain := domainFromModel(spaceModel)

	return spaceDomain, nil
}
