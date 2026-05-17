package tasks_http_transport

import (
	"time"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
)

type TaskDTOResponse struct {
	ID          uuid.UUID `json:"id"`
	ListID      uuid.UUID `json:"list_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Status      string    `json:"status"`
	Position    int       `json:"position"`
	CreatedAt   time.Time `json:"created_at"`
}

func listDTOResponseFromDomain(taskDomain domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:          taskDomain.ID,
		ListID:      taskDomain.ListID,
		Title:       taskDomain.Title,
		Description: taskDomain.Description,
		Status:      taskDomain.Status,
		Position:    taskDomain.Position,
		CreatedAt:   taskDomain.CreatedAt,
	}
}

func listDTOResponsesFromDomains(taskDomains []domain.Task) []TaskDTOResponse {
	dtos := make([]TaskDTOResponse, len(taskDomains))
	for i, taskDomain := range taskDomains {
		dtos[i] = listDTOResponseFromDomain(taskDomain)
	}

	return dtos
}
