package auth_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	auth_service "github.com/sqlmerr/huddle/backend/internal/auth/service"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_request "github.com/sqlmerr/huddle/backend/internal/core/transport/http/request"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	"go.uber.org/zap"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1"`
}

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *AuthHTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	input := serviceDTOFromDTO(request)
	user, err := h.authService.Register(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(fmt.Errorf("create user: %w", err), "failed to create user")
		return
	}

	log.Debug("registered new user", zap.String("user_id", user.ID.String()))
	responseHandler.JSONResponse(http.StatusCreated, CreateUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
}

func serviceDTOFromDTO(request CreateUserRequest) auth_service.RegisterInput {
	return auth_service.RegisterInput{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
}
