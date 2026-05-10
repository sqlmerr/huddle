package spaces_http_transport

import (
	"time"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type SpaceDTOResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	OwnerID     uuid.UUID `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func spaceDTOResponseFromDomain(spaceDomain domain.Space) SpaceDTOResponse {
	return SpaceDTOResponse{
		ID:          spaceDomain.ID,
		Title:       spaceDomain.Title,
		Description: spaceDomain.Description,
		OwnerID:     spaceDomain.OwnerID,
		CreatedAt:   spaceDomain.CreatedAt,
	}
}

func spaceDTOResponsesFromDomains(spaceDomains []domain.Space) []SpaceDTOResponse {
	dtos := make([]SpaceDTOResponse, len(spaceDomains))
	for i, spaceDomain := range spaceDomains {
		dtos[i] = spaceDTOResponseFromDomain(spaceDomain)
	}

	return dtos
}
