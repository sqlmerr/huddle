package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (r *TaskRepositoryImpl) DeleteTask(ctx context.Context, taskID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM tasks
	WHERE id = $1;
	`

	cmtTag, err := r.pool.Exec(ctx, query, taskID)
	if err != nil {
		return fmt.Errorf("exec error: %w", err)
	}
	if cmtTag.RowsAffected() == 0 {
		return fmt.Errorf("task with id='%s': %w", taskID, core_errors.ErrNotFound)
	}

	return nil
}
