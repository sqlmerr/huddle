package tasks_http_transport

import (
	"net/http"

	"github.com/google/uuid"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
	tasks_service "github.com/sqlmerr/huddle/backend/internal/tasks/service"
)

type MoveTaskRequest struct {
	ListID   uuid.UUID `json:"list_id" validate:"required"`
	Position *int      `json:"position" validate:"omitempty,min=0"`
}

type MoveTaskResponse TaskDTOResponse

func (h *TaskHTTPHandler) MoveTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request MoveTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)
	taskID, err := core_http_utils.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get `id` path value")
		return
	}

	input := tasks_service.MoveTaskInput{
		TaskID:   taskID,
		ListID:   request.ListID,
		Position: request.Position,
	}
	task, err := h.taskService.MoveTask(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to move task")
		return
	}

	response := MoveTaskResponse(listDTOResponseFromDomain(task))
	responseHandler.JSONResponse(http.StatusOK, response)
}
