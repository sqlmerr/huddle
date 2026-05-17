package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (r *TaskRepositoryImpl) GetListTasks(ctx context.Context, listID uuid.UUID) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, list_id, title, description, status, position, created_at
	FROM tasks
	WHERE list_id = $1
	ORDER BY position;
	`

	rows, err := r.pool.Query(ctx, query, listID)
	if err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	var taskModels []TaskModel
	for rows.Next() {
		var m TaskModel
		err := rows.Scan(
			&m.ID,
			&m.ListID,
			&m.Title,
			&m.Description,
			&m.Status,
			&m.Position,
			&m.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		taskModels = append(taskModels, m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	tasks := domainFromModels(taskModels)
	return tasks, nil
}
