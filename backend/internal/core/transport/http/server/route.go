package core_http_server

import "net/http"

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func NewRoute(method, path string, handler http.HandlerFunc) *Route {
	return &Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}
