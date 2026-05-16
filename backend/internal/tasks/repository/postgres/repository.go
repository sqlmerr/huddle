package tasks_postgres_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTask(ctx context.Context, taskID uuid.UUID) (domain.Task, error)
	GetListTasks(ctx context.Context, listID uuid.UUID) ([]domain.Task, error)
	SaveTask(ctx context.Context, task domain.Task) (domain.Task, error)
	DeleteTask(ctx context.Context, taskID uuid.UUID) error
}

type TaskRepositoryImpl struct {
	pool core_postgres_pool.Pool
}

func NewTaskRepository(pool core_postgres_pool.Pool) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{pool}
}
