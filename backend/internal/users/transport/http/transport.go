package users_transport_http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService   UsersService
	authMiddleware core_http_middleware.Middleware
}

type UsersService interface {
	GetUser(ctx context.Context, userID uuid.UUID) (domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService, authMiddleware core_http_middleware.Middleware) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService:   usersService,
		authMiddleware: authMiddleware,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodGet,
			Path:       "/users/me",
			Handler:    h.GetMe,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
	}
}
