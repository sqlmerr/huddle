package users_transport_http

import (
	"net/http"

	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
	users_service "github.com/sqlmerr/huddle/backend/internal/users/service"
)

type UserHTTPHandler struct {
	usersService   users_service.UserService
	authMiddleware core_http_middleware.Middleware
}

func NewUserHTTPHandler(usersService users_service.UserService, authMiddleware core_http_middleware.Middleware) *UserHTTPHandler {
	return &UserHTTPHandler{
		usersService:   usersService,
		authMiddleware: authMiddleware,
	}
}

func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodGet,
			Path:       "/users/me",
			Handler:    h.GetMe,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
	}
}
