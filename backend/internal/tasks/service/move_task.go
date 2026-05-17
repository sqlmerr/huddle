package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

func (s *TaskServiceImpl) MoveTask(ctx context.Context, userID uuid.UUID, input MoveTaskInput) (domain.Task, error) {
	task, err := s.repo.GetTask(ctx, input.TaskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}

	if err := s.accessService.CanAccessTask(ctx, userID, task); err != nil {
		return domain.Task{}, fmt.Errorf("can't access task: %w", err)
	}

	if input.ListID == task.ListID {
		return domain.Task{}, fmt.Errorf(
			"can't move task to the same list. Use ReorderTasks method instead: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if err := s.accessService.CanAccessListByID(ctx, userID, input.ListID); err != nil { // TODO: Permissions. CanAccessListByID(ctx, userID, listID, permissions.PermissionCreateTask) or smth
		return domain.Task{}, fmt.Errorf("can't access list: %w", err)
	}

	currentListTasks, err := s.repo.GetListTasks(ctx, task.ListID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get current list tasks: %w", err)
	}

	newListTasks, err := s.repo.GetListTasks(ctx, input.ListID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get new list tasks: %w", err)
	}

	newTask := task
	newTask.ListID = input.ListID

	if input.Position != nil {
		if *input.Position < 0 || *input.Position > len(newListTasks) {
			return domain.Task{}, fmt.Errorf(
				"invalid position: %d (must be between 0 and %d): %w",
				*input.Position,
				len(newListTasks),
				core_errors.ErrInvalidArgument,
			)
		}

		newTask.Position = *input.Position
	} else {
		newTask.Position = len(newListTasks)
	}

	for i := range newListTasks {
		if newListTasks[i].Position >= newTask.Position {
			newListTasks[i].Position++
			_, err := s.repo.SaveTask(ctx, newListTasks[i])
			if err != nil {
				return domain.Task{}, fmt.Errorf("update new list task position: %w", err)
			}
		}
	}

	for i := range currentListTasks {
		if currentListTasks[i].ID == task.ID {
			continue
		}
		if currentListTasks[i].Position > task.Position {
			currentListTasks[i].Position--
			_, err := s.repo.SaveTask(ctx, currentListTasks[i])
			if err != nil {
				return domain.Task{}, fmt.Errorf("update old list task position: %w", err)
			}
		}
	}

	updatedTask, err := s.repo.SaveTask(ctx, newTask)
	if err != nil {
		return domain.Task{}, fmt.Errorf("save moved task: %w", err)
	}

	return updatedTask, nil
}
