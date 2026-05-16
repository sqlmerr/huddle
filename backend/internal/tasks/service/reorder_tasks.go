package tasks_service

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (s *TaskServiceImpl) ReorderTasks(ctx context.Context, userID uuid.UUID, input ReorderTasksInput) error {
	if err := s.accessService.CanAccessListByID(ctx, userID, input.ListID); err != nil {
		return fmt.Errorf("can't access list: %w", err)
	}

	tasks, err := s.repo.GetListTasks(ctx, input.ListID)
	if err != nil {
		return fmt.Errorf("get list tasks: %w", err)
	}

	if len(input.Order) != len(tasks) {
		return fmt.Errorf(
			"`Order` must contain ALL task ids ONLY for this list: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	newTasks := make([]domain.Task, len(tasks))
	for i, task := range tasks {
		if !slices.Contains(input.Order, task.ID) {
			return fmt.Errorf(
				"`Order` must contain ALL task ids for this list: %w",
				core_errors.ErrInvalidArgument,
			)
		}
		newTask := task
		newTask.Position = slices.Index(input.Order, task.ID)
		newTasks[i] = newTask
	}

	for _, newTask := range newTasks {
		_, err = s.repo.SaveTask(ctx, newTask)
		if err != nil {
			return fmt.Errorf("save task: %w", err)
		}
	}

	return nil
}
