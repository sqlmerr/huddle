package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *TaskServiceImpl) CreateTask(ctx context.Context, userID uuid.UUID, input CreateTaskInput) (domain.Task, error) {
	if err := s.accessService.CanAccessListByID(ctx, userID, input.ListID); err != nil {
		return domain.Task{}, fmt.Errorf("cannot access list: %w", err)
	}

	tasks, err := s.repo.GetListTasks(ctx, input.ListID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get list tasks: %w", err)
	}

	var position int
	if len(tasks) != 0 {
		position = tasks[len(tasks)-1].Position + 1
	}

	task := domain.NewTaskUninitialized(
		input.ListID,
		input.Title,
		input.Description,
		input.Status,
		position,
	)
	taskDomain, err := s.repo.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("create task: %w", err)
	}
	return taskDomain, nil
}
