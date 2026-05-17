package tasks_http_transport

import (
	"net/http"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
)

type GetListTasksResponse struct {
	Data []TaskDTOResponse `json:"data"`
}

func (h *TaskHTTPHandler) GetListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID := core_auth.GetUserIDFromContext(ctx)
	listID, err := core_http_utils.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get `id` path value")
		return
	}

	tasks, err := h.taskService.GetListTasks(ctx, userID, listID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get list tasks")
		return
	}

	response := GetListTasksResponse{Data: listDTOResponsesFromDomains(tasks)}
	responseHandler.JSONResponse(http.StatusOK, response)
}
