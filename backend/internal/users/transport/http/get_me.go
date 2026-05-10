package users_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
)

type GetMeResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *UserHTTPHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID := core_auth.GetUserIDFromContext(ctx)

	user, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(fmt.Errorf("%w: %w", core_errors.ErrInternalServerError, err), "failed to get user")
	}

	responseHandler.JSONResponse(http.StatusOK, GetMeResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
}
