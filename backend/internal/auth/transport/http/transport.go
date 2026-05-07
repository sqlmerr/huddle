package auth_transport_http

import (
	"context"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/domain"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
)

type AuthHTTPHandler struct {
	authService AuthService
}

type AuthService interface {
	Register(ctx context.Context, user domain.User) (domain.User, error)
	LoginByUsername(ctx context.Context, username, password string) (core_auth.Token, error)
	LoginByEmail(ctx context.Context, email, password string) (core_auth.Token, error)
}

func NewAuthHTTPHandler(authService AuthService) *AuthHTTPHandler {
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
