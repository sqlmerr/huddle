package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

func (s *TaskServiceImpl) GetListTasks(ctx context.Context, userID uuid.UUID, listID uuid.UUID) ([]domain.Task, error) {
	if err := s.accessService.CanAccessListByID(ctx, userID, listID); err != nil {
		return nil, fmt.Errorf("can't access list: %w", err)
	}

	tasks, err := s.repo.GetListTasks(ctx, listID)
	if err != nil {
		return nil, fmt.Errorf("get list tasks: %w", err)
	}

	return tasks, nil
}
