package spaces_http_transport

import (
	"net/http"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
)

type GetMySpacesResponse struct {
	Data []SpaceDTOResponse `json:"data"`
}

func (h *SpaceHTTPHandler) GetMySpaces(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID := core_auth.GetUserIDFromContext(ctx)
	spaces, err := h.spaceService.GetUserSpaces(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user's spaces")
		return
	}
	response := GetMySpacesResponse{spaceDTOResponsesFromDomains(spaces)}
	responseHandler.JSONResponse(http.StatusOK, response)
}
