package core_http_server

import (
	"fmt"
	"net/http"
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
		var handler http.Handler = http.HandlerFunc(route.Handler)
		for _, m := range route.Middleware {
			handler = m(handler)
		}
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.Handle(pattern, handler)
	}
}
