package boards_http_transport

import (
	"net/http"

	"github.com/google/uuid"
	boards_service "github.com/sqlmerr/huddle/backend/internal/boards/service"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
)

type CreateBoardRequest struct {
	Title   string    `json:"title" validate:"required,min=1,max=50"`
	SpaceID uuid.UUID `json:"space_id" validate:"required"`
}

type CreateBoardResponse BoardDTOResponse

func (h *BoardsHTTPHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request CreateBoardRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and valdate HTTP request body")
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)

	input := boards_service.CreateBoardInput{
		Title:   request.Title,
		SpaceID: request.SpaceID,
	}
	board, err := h.boardService.CreateBoard(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create board")
		return
	}

	response := CreateBoardResponse(boardDTOResponseFromDomain(board))
	responseHandler.JSONResponse(http.StatusCreated, response)
}
