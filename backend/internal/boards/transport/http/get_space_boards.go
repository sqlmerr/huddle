package boards_http_transport

import (
	"net/http"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
)

type GetSpaceBoardsResponse struct {
	Data []BoardDTOResponse `json:"data"`
}

func (h *BoardsHTTPHandler) GetSpaceBoards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	spaceID, err := core_http_utils.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get `id` path value")
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)

	boards, err := h.boardService.GetSpaceBoards(ctx, userID, spaceID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get space's boards")
		return
	}

	response := GetSpaceBoardsResponse{
		Data: boardDTOResponsesFromDomains(boards),
	}
	responseHandler.JSONResponse(http.StatusOK, response)
}
