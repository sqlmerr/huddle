package lists_http_transport

import (
	"fmt"
	"net/http"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	core_http_types "github.com/sqlmerr/huddle/backend/internal/core/transport/http/types"
	core_http_utils "github.com/sqlmerr/huddle/backend/internal/core/transport/http/utils"
	lists_service "github.com/sqlmerr/huddle/backend/internal/lists/service"
)

type PatchListRequest struct {
	Title core_http_types.Nullable[string] `json:"title"`
}

type PatchListResponse ListDTOResponse

func (r *PatchListRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("`Title` can't be null")
		}
		titleLength := len([]rune(*r.Title.Value))
		if titleLength == 0 || titleLength > 50 {
			return fmt.Errorf("`Title` length must be between 1 and 50 characters")
		}
	}

	return nil
}

func (h *ListHTTPHandler) PatchList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request PatchListRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)
	listID, err := core_http_utils.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get `id` path value")
		return
	}

	input := lists_service.PatchListInput{
		ListID: listID,
		Title:  request.Title.ToDomain(),
	}
	list, err := h.listService.PatchList(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch list")
		return
	}

	response := PatchListResponse(listDTOResponseFromDomain(list))
	responseHandler.JSONResponse(http.StatusOK, response)
}
