package lists_http_transport

import (
	"net/http"

	"github.com/google/uuid"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
	lists_service "github.com/sqlmerr/huddle/backend/internal/lists/service"
)

type ReorderListsRequest struct {
	Order []uuid.UUID `json:"order" validate:"required"`
}

func (h *ListHTTPHandler) ReorderLists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request ReorderListsRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)
	boardID, err := core_http_utils.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get `id` path value")
		return
	}

	input := lists_service.ReorderListsInput{
		BoardID: boardID,
		Order:   request.Order,
	}
	err = h.listService.ReorderLists(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to reorder lists")
		return
	}

	responseHandler.NoContentResponse()
}
