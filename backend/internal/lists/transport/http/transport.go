package lists_http_transport

import (
	"net/http"

	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
	lists_service "github.com/sqlmerr/huddle/backend/internal/lists/service"
)

type ListHTTPHandler struct {
	listService    lists_service.ListService
	authMiddleware core_http_middleware.Middleware
}

func NewListHTTPHandler(
	listService lists_service.ListService,
	authMiddleware core_http_middleware.Middleware,
) *ListHTTPHandler {
	return &ListHTTPHandler{
		listService, authMiddleware,
	}
}

func (h *ListHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodPost,
			Path:       "/lists",
			Handler:    h.CreateList,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/lists/{id}",
			Handler:    h.GetList,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/boards/{id}/lists",
			Handler:    h.GetBoardLists,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodPatch,
			Path:       "/lists/{id}",
			Handler:    h.PatchList,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodDelete,
			Path:       "/lists/{id}",
			Handler:    h.DeleteList,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodPatch,
			Path:       "/boards/{id}/lists/reorder",
			Handler:    h.ReorderLists,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
	}
}
