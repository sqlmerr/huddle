package auth_transport_http

import (
	"net/http"

	auth_service "github.com/sqlmerr/huddle/backend/internal/auth/service"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
)

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=1"`
}

func (h *AuthHTTPHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request ChangePasswordRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	userID := core_auth.GetUserIDFromContext(ctx)

	input := auth_service.ChangePasswordInput{UserID: userID, OldPassword: request.OldPassword, NewPassword: request.NewPassword}
	err := h.authService.ChangePassword(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to change user's password")
		return
	}

	responseHandler.NoContentResponse()
}
