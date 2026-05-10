package spaces_http_transport

import (
	"fmt"
	"net/http"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_types "github.com/sqlmerr/huddle/backend/internal/core/transport/http/types"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
	spaces_service "github.com/sqlmerr/huddle/backend/internal/spaces/service"
)

type PatchSpaceRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	IsArchived  core_http_types.Nullable[bool]   `json:"is_archived"`
}

type PatchSpaceResponse SpaceDTOResponse

func (r *PatchSpaceRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("`Title` can't be null")
		}
		titleLength := len([]rune(*r.Title.Value))
		if titleLength < 1 || titleLength > 50 {
			return fmt.Errorf("`Title` length must be between 1 and 50 characters")
		}
	}

	if r.Description.Set && r.Description.Value != nil {
		descriptionLength := len([]rune(*r.Description.Value))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf("`Description` length must be between 1 and 1000 characters")
		}
	}

	if r.IsArchived.Set && r.IsArchived.Value == nil {
		return fmt.Errorf("`IsArchived` can't be null")
	}

	return nil
}

func (h *SpaceHTTPHandler) PatchSpace(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request PatchSpaceRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request body",
		)
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)
	spaceID, err := core_http_utils.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get `id` path value")
		return
	}

	input := spaces_service.PatchSpaceInput{
		UserID:      userID,
		SpaceID:     spaceID,
		Title:       request.Title.ToDomain(),
		Description: request.Description.ToDomain(),
		IsArchived:  request.IsArchived.ToDomain(),
	}
	space, err := h.spaceService.PatchSpace(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch space")
		return
	}

	response := PatchSpaceResponse(spaceDTOResponseFromDomain(space))
	responseHandler.JSONResponse(http.StatusOK, response)
}
