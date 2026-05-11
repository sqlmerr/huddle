package boards_http_transport

import (
	"net/http"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
)

func (h *BoardsHTTPHandler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID := core_auth.GetUserIDFromContext(ctx)
	boardID, err := core_http_utils.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get `id` path value")
		return
	}

	err = h.boardService.DeleteBoard(ctx, userID, boardID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to delete board")
		return
	}

	responseHandler.NoContentResponse()
}
