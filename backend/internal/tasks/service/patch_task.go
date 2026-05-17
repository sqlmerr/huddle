package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *TaskServiceImpl) PatchTask(ctx context.Context, userID uuid.UUID, input PatchTaskInput) (domain.Task, error) {
	task, err := s.repo.GetTask(ctx, input.TaskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}

	if err := s.accessService.CanAccessTask(ctx, userID, task); err != nil {
		return domain.Task{}, fmt.Errorf("can't access task: %w", err)
	}

	patchedTask, err := input.ApplyPatch(task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("apply patch: %w", err)
	}

	taskDomain, err := s.repo.SaveTask(ctx, patchedTask)
	if err != nil {
		return domain.Task{}, fmt.Errorf("save task: %w", err)
	}

	return taskDomain, nil
}
