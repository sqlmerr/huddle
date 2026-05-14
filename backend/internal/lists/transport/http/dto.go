package lists_http_transport

import (
	"time"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type ListDTOResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	BoardID   uuid.UUID `json:"board_id"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at"`
}

func listDTOResponseFromDomain(listDomain domain.List) ListDTOResponse {
	return ListDTOResponse{
		ID:        listDomain.ID,
		Title:     listDomain.Title,
		BoardID:   listDomain.BoardID,
		Position:  listDomain.Position,
		CreatedAt: listDomain.CreatedAt,
	}
}

func listDTOResponsesFromDomains(listDomains []domain.List) []ListDTOResponse {
	dtos := make([]ListDTOResponse, len(listDomains))
	for i, listDomain := range listDomains {
		dtos[i] = listDTOResponseFromDomain(listDomain)
	}

	return dtos
}
