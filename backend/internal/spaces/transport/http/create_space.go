package spaces_http_transport

import (
	"net/http"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	spaces_service "github.com/sqlmerr/huddle/backend/internal/spaces/service"
	"go.uber.org/zap"
)

type CreateSpaceRequest struct {
	Title       string  `json:"title" validate:"required,min=1,max=50"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000"`
}

type CreateSpaceResponse SpaceDTOResponse

func (h *SpaceHTTPHandler) CreateSpace(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request CreateSpaceRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)
	input := spaces_service.CreateSpaceInput{
		OwnerID:     userID,
		Title:       request.Title,
		Description: request.Description,
	}
	space, err := h.spaceService.CreateSpace(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create space")
		return
	}

	log.Debug(
		"created space",
		zap.String("space_id", space.ID.String()),
		zap.String("owner_id", space.OwnerID.String()),
	)

	response := CreateSpaceResponse(spaceDTOResponseFromDomain(space))
	responseHandler.JSONResponse(http.StatusCreated, response)
}
