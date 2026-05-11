package boards_http_transport

import (
	"net/http"

	boards_service "github.com/sqlmerr/huddle/backend/internal/boards/service"
	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
)

type BoardsHTTPHandler struct {
	boardService   boards_service.BoardService
	authMiddleware core_http_middleware.Middleware
}

func NewBoardsHTTPHandler(
	boardService boards_service.BoardService,
	authMiddleware core_http_middleware.Middleware,
) *BoardsHTTPHandler {
	return &BoardsHTTPHandler{boardService, authMiddleware}
}

func (h *BoardsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodPost,
			Path:       "/boards",
			Handler:    h.CreateBoard,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/spaces/{id}/boards",
			Handler:    h.GetSpaceBoards,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/boards/{id}",
			Handler:    h.GetBoard,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodPatch,
			Path:       "/boards/{id}",
			Handler:    h.PatchBoard,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodDelete,
			Path:       "/boards/{id}",
			Handler:    h.DeleteBoard,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
	}
}
