package spaces_http_transport

import (
	"net/http"

	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
	spaces_service "github.com/sqlmerr/huddle/backend/internal/spaces/service"
)

type SpaceHTTPHandler struct {
	spaceService   spaces_service.SpaceService
	authMiddleware core_http_middleware.Middleware
}

func NewSpaceHTTPHandler(spaceService spaces_service.SpaceService, authMiddleware core_http_middleware.Middleware) *SpaceHTTPHandler {
	return &SpaceHTTPHandler{
		spaceService:   spaceService,
		authMiddleware: authMiddleware,
	}
}

func (h *SpaceHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodPost,
			Path:       "/spaces",
			Handler:    h.CreateSpace,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/spaces/my",
			Handler:    h.GetMySpaces,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/spaces/{id}",
			Handler:    h.GetSpace,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodPatch,
			Path:       "/spaces/{id}",
			Handler:    h.PatchSpace,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
	}
}
