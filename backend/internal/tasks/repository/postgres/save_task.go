package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *TaskRepositoryImpl) SaveTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE tasks
	SET
		list_id = $2,
		title = $3,
		description = $4,
		status = $5,
		position = $6
	WHERE id = $1
	RETURNING id, list_id, title, description, status, position, created_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		task.ID,
		task.ListID,
		task.Title,
		task.Description,
		task.Status,
		task.Position,
	)

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
				task.ID,
				core_errors.ErrNotFound,
			)
		}

		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf(
				"list with id='%s': %w",
				task.ListID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := domainFromModel(m)
	return taskDomain, nil
}
