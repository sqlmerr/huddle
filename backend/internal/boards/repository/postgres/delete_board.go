package boards_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (r *BoardRepositoryImpl) DeleteBoard(ctx context.Context, boardID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM boards
	WHERE id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, boardID)
	if err != nil {
		return fmt.Errorf("exec error: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("board with id='%s': %w", boardID, core_errors.ErrNotFound)
	}

	return nil
}
