package users_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	"go.uber.org/zap"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	// ...

	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error("decode request body", zap.Error(err))
		responseHandler.ErrorResponse(http.StatusBadRequest, err, "decode request body")
		return
	}

	responseHandler.JSONResponse(http.StatusCreated, map[string]bool{"OK": true})
}
