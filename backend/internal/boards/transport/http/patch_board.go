package boards_http_transport

import (
	"fmt"
	"net/http"

	boards_service "github.com/sqlmerr/huddle/backend/internal/boards/service"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_types "github.com/sqlmerr/huddle/backend/internal/core/transport/http/types"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
)

type PatchBoardRequest struct {
	Title core_http_types.Nullable[string]
}

type PatchBoardResponse BoardDTOResponse

func (r *PatchBoardRequest) Validate() error {
	if r.Title.Set && r.Title.Value == nil {
		return fmt.Errorf("`Title` can't be null")
	}

	return nil
}

func (h *BoardsHTTPHandler) PatchBoard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request PatchBoardRequest
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

	input := boards_service.PatchBoardInput{
		BoardID: boardID,
		Title:   request.Title.ToDomain(),
	}
	board, err := h.boardService.PatchBoard(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch board")
		return
	}

	response := PatchBoardResponse(boardDTOResponseFromDomain(board))
	responseHandler.JSONResponse(http.StatusOK, response)
}
