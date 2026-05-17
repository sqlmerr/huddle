package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *TaskServiceImpl) DeleteTask(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) error {
	if err := s.accessService.CanAccessTaskByID(ctx, userID, taskID); err != nil {
		return fmt.Errorf("can't access task: %w", err)
	}

	err := s.repo.DeleteTask(ctx, taskID)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}
	return nil
}
