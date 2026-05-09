package core_http_server

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
)

type ApiVersion string

const (
	ApiVersionV1 ApiVersion = "v1"
	ApiVersionV2 ApiVersion = "v2"
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
}

func NewAPIVersionRouter(apiVersion ApiVersion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
	}
}

func (r *APIVersionRouter) AddRoutes(routes ...Route) {
	for _, route := range routes {
		handler := core_http_middleware.ChainMiddleware(route.Handler, route.Middleware...)
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.Handle(pattern, handler)
	}
}
