package auth_transport_http

import (
	auth_service "github.com/sqlmerr/huddle/backend/internal/auth/service"
	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
)

type AuthHTTPHandler struct {
	authService    auth_service.AuthService
	authMiddleware core_http_middleware.Middleware
}

func NewAuthHTTPHandler(authService auth_service.AuthService, authMiddleware core_http_middleware.Middleware) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService:    authService,
		authMiddleware: authMiddleware,
	}
}

func (h *AuthHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Path:    "/auth/register",
			Method:  "POST",
			Handler: h.Register,
		},
		{
			Path:    "/auth/login",
			Method:  "POST",
			Handler: h.Login,
		},
		{
			Path:       "/auth/password",
			Method:     "PATCH",
			Handler:    h.ChangePassword,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
	}
}
