package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
)

type Route struct {
	Path       string
	Method     string
	Handler    http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}

func NewRoute(method, path string, handler http.HandlerFunc, middleware ...core_http_middleware.Middleware) *Route {
	return &Route{
		Path:       path,
		Method:     method,
		Handler:    handler,
		Middleware: middleware,
	}
}
