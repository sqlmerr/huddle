package users_transport_http

import (
	"net/http"

	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
}

func NewUsersHTTPHandler() *UsersHTTPHandler {
	return &UsersHTTPHandler{}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
