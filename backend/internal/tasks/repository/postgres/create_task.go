package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

func (r *TaskRepositoryImpl) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO tasks (list_id, title, description, status, position)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, list_id, title, description, status, position, created_at;
	`

	row := r.pool.QueryRow(ctx, query, task.ListID, task.Title, task.Description, task.Status, task.Position)

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
