package lists_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (r *ListRepositoryImpl) DeleteList(ctx context.Context, listID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM lists
	WHERE id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, listID)

	if err != nil {
		return fmt.Errorf("exec error: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("list with id='%s': %w", listID, core_errors.ErrNotFound)
	}

	return nil
}
