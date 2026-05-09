package auth_transport_http

import (
	auth_service "github.com/sqlmerr/huddle/backend/internal/auth/service"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
)

type AuthHTTPHandler struct {
	authService auth_service.AuthService
}

func NewAuthHTTPHandler(authService auth_service.AuthService) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
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
	}
}
