package lists_postgres_repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type ListModel struct {
	ID        uuid.UUID
	Title     string
	BoardID   uuid.UUID
	Position  int
	CreatedAt time.Time
}

func domainFromModels(models []ListModel) []domain.List {
	domains := make([]domain.List, len(models))
	for i, model := range models {
		domains[i] = domainFromModel(model)
	}
	return domains
}

func domainFromModel(model ListModel) domain.List {
	return domain.NewList(
		model.ID,
		model.Title,
		model.BoardID,
		model.Position,
		model.CreatedAt,
	)
}
