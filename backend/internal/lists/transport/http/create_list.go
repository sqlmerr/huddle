package lists_http_transport

import (
	"net/http"

	"github.com/google/uuid"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	lists_service "github.com/sqlmerr/huddle/backend/internal/lists/service"
	"go.uber.org/zap"
)

type CreateListRequest struct {
	Title   string    `json:"title" validate:"required,min=1,max=50"`
	BoardID uuid.UUID `json:"board_id" validate:"required"`
}

type CreateListResponse ListDTOResponse

func (h *ListHTTPHandler) CreateList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request CreateListRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)

	input := lists_service.CreateListInput{
		Title:   request.Title,
		BoardID: request.BoardID,
	}
	list, err := h.listService.CreateList(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create list")
		return
	}

	log.Debug(
		"created task list",
		zap.String("list_id", list.ID.String()),
		zap.String("board_id", list.BoardID.String()),
	)

	response := CreateListResponse(listDTOResponseFromDomain(list))
	responseHandler.JSONResponse(http.StatusCreated, response)
}
