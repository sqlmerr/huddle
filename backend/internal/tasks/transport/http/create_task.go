package tasks_http_transport

import (
	"net/http"

	"github.com/google/uuid"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	tasks_service "github.com/sqlmerr/huddle/backend/internal/tasks/service"
)

type CreateTaskRequest struct {
	ListID      uuid.UUID `json:"list_id" validate:"required"`
	Title       string    `json:"title" validate:"required,min=1,max=255"`
	Description *string   `json:"description" validate:"omitempty,min=1,max=1000"`
	Status      string    `json:"status" validate:"required,min=1,max=50"`
}

type CreateTaskResponse TaskDTOResponse

func (h *TaskHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request CreateTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)
	input := tasks_service.CreateTaskInput{
		ListID:      request.ListID,
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	}

	task, err := h.taskService.CreateTask(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task")
		return
	}

	response := CreateTaskResponse(listDTOResponseFromDomain(task))
	responseHandler.JSONResponse(http.StatusCreated, response)
}
