package auth_transport_http

import (
	"fmt"
	"net/http"

	auth_service "github.com/sqlmerr/huddle/backend/internal/auth/service"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
)

type LoginRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email" validate:"omitempty,email"`
	Password string  `json:"password" validate:"required,min=1"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (h *AuthHTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request LoginRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	if (request.Username == nil || *request.Username == "") && (request.Email == nil || *request.Email == "") {
		responseHandler.ErrorResponse(fmt.Errorf("`Username` or `Email` must be non-null: %w", core_errors.ErrInvalidArgument), "at least one of the username or email must be given.")
		return
	}

	var token core_auth.Token
	var err error
	if request.Username != nil {
		token, err = h.authService.LoginByUsername(ctx, auth_service.LoginByUsernameInput{Username: *request.Username, Password: request.Password})
	} else {
		token, err = h.authService.LoginByEmail(ctx, auth_service.LoginByEmailInput{Email: *request.Email, Password: request.Password})
	}

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to login")
		return
	}

	log.Debug("user successfully logged in")
	response := LoginResponse{AccessToken: token.AccessToken, TokenType: "Bearer"}
	responseHandler.JSONResponse(http.StatusOK, response)
}
