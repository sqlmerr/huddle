package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *TaskServiceImpl) GetTask(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) (domain.Task, error) {
	task, err := s.repo.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}

	if err := s.accessService.CanAccessTask(ctx, userID, task); err != nil {
		return domain.Task{}, fmt.Errorf("can't accesss task: %w", err)
	}

	return task, nil
}
