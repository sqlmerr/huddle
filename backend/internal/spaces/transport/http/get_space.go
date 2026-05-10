package spaces_http_transport

import (
	"net/http"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
)

type GetSpaceResponse SpaceDTOResponse

func (h *SpaceHTTPHandler) GetSpace(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID := core_auth.GetUserIDFromContext(ctx)
	spaceID, err := core_http_utils.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get `id` path value")
		return
	}
	space, err := h.spaceService.GetSpace(ctx, userID, spaceID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get space")
		return
	}

	response := spaceDTOResponseFromDomain(space)
	responseHandler.JSONResponse(http.StatusOK, response)
}
