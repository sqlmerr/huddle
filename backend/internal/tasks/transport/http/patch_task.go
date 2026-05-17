package tasks_http_transport

import (
	"fmt"
	"net/http"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_types "github.com/sqlmerr/huddle/backend/internal/core/transport/http/types"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
	tasks_service "github.com/sqlmerr/huddle/backend/internal/tasks/service"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Status      core_http_types.Nullable[string] `json:"status"`
}

type PatchTaskResponse TaskDTOResponse

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set && r.Title.Value == nil {
		return fmt.Errorf("`Title` can't be null: %w", core_errors.ErrInvalidArgument)
	}

	if r.Status.Set && r.Status.Value == nil {
		return fmt.Errorf("`Status` can't be null: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (h *TaskHTTPHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request PatchTaskRequest
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
	input := tasks_service.PatchTaskInput{
		TaskID:      taskID,
		Title:       request.Title.ToDomain(),
		Description: request.Description.ToDomain(),
		Status:      request.Status.ToDomain(),
	}

	task, err := h.taskService.PatchTask(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch task")
		return
	}

	response := PatchTaskResponse(listDTOResponseFromDomain(task))
	responseHandler.JSONResponse(http.StatusOK, response)
}
