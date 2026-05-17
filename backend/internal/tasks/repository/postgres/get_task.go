package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *TaskRepositoryImpl) GetTask(ctx context.Context, taskID uuid.UUID) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, list_id, title, description, status, position, created_at
	FROM tasks
	WHERE id = $1;
	`

	row := r.pool.QueryRow(ctx, query, taskID)

	var m TaskModel
	err := row.Scan(
		&m.ID,
		&m.ListID,
		&m.Title,
		&m.Description,
		&m.Status,
		&m.Position,
		&m.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%s': %w",
				taskID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := domainFromModel(m)
	return taskDomain, nil
}
