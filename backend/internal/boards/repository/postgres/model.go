package boards_postgres_repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type BoardModel struct {
	ID        uuid.UUID
	Title     string
	SpaceID   uuid.UUID
	CreatedAt time.Time
}

func domainFromModels(models []BoardModel) []domain.Board {
	domains := make([]domain.Board, len(models))
	for i, model := range models {
		domains[i] = domainFromModel(model)
	}
	return domains
}

func domainFromModel(model BoardModel) domain.Board {
	return domain.NewBoard(
		model.ID,
		model.Title,
		model.SpaceID,
		model.CreatedAt,
	)
}
