package tasks_service

import (
	"context"

	"github.com/google/uuid"
	core_access "github.com/sqlmerr/huddle/backend/internal/core/access"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	tasks_postgres_repository "github.com/sqlmerr/huddle/backend/internal/tasks/repository/postgres"
)

type TaskService interface {
	CreateTask(ctx context.Context, userID uuid.UUID, input CreateTaskInput) (domain.Task, error)
	GetTask(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) (domain.Task, error)
	GetListTasks(ctx context.Context, userID uuid.UUID, listID uuid.UUID) ([]domain.Task, error)
	PatchTask(ctx context.Context, userID uuid.UUID, input PatchTaskInput) (domain.Task, error)
	DeleteTask(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) error
	ReorderTasks(ctx context.Context, userID uuid.UUID, input ReorderTasksInput) error
	MoveTask(ctx context.Context, userID uuid.UUID, input MoveTaskInput) (domain.Task, error)
}

type TaskServiceImpl struct {
	repo          tasks_postgres_repository.TaskRepository
	accessService core_access.AccessService
}

func NewTaskService(
	repo tasks_postgres_repository.TaskRepository,
	accessService core_access.AccessService,
) *TaskServiceImpl {
	return &TaskServiceImpl{repo, accessService}
}
