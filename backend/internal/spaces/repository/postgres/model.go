package spaces_postgres_repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type SpaceModel struct {
	ID          uuid.UUID
	Title       string
	Description *string
	OwnerID     uuid.UUID
	CreatedAt   time.Time
}

func domainFromModels(models []SpaceModel) []domain.Space {
	domains := make([]domain.Space, len(models))
	for i, model := range models {
		domains[i] = domainFromModel(model)
	}
	return domains
}

func domainFromModel(model SpaceModel) domain.Space {
	return domain.NewSpace(
		model.ID,
		model.Title,
		model.Description,
		model.OwnerID,
		model.CreatedAt,
	)
}
