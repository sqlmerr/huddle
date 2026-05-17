package tasks_postgres_repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type TaskModel struct {
	ID          uuid.UUID
	ListID      uuid.UUID
	Title       string
	Description *string
	Status      string
	Position    int
	CreatedAt   time.Time
}

func domainFromModels(models []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(models))
	for i, model := range models {
		domains[i] = domainFromModel(model)
	}
	return domains
}

func domainFromModel(model TaskModel) domain.Task {
	return domain.NewTask(
		model.ID,
		model.ListID,
		model.Title,
		model.Description,
		model.Status,
		model.Position,
		model.CreatedAt,
	)
}
