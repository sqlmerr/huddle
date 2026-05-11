package boards_http_transport

import (
	"time"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type BoardDTOResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	SpaceID   uuid.UUID `json:"space_id"`
	CreatedAt time.Time `json:"created_at"`
}

func boardDTOResponseFromDomain(boardDomain domain.Board) BoardDTOResponse {
	return BoardDTOResponse{
		ID:        boardDomain.ID,
		Title:     boardDomain.Title,
		SpaceID:   boardDomain.SpaceID,
		CreatedAt: boardDomain.CreatedAt,
	}
}

func boardDTOResponsesFromDomains(boardDomains []domain.Board) []BoardDTOResponse {
	dtos := make([]BoardDTOResponse, len(boardDomains))
	for i, boardDomain := range boardDomains {
		dtos[i] = boardDTOResponseFromDomain(boardDomain)
	}

	return dtos
}
