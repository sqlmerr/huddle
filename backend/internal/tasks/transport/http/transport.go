package tasks_http_transport

import (
	"net/http"

	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
	tasks_service "github.com/sqlmerr/huddle/backend/internal/tasks/service"
)

type TaskHTTPHandler struct {
	taskService    tasks_service.TaskService
	authMiddleware core_http_middleware.Middleware
}

func NewTaskHTTPHandler(
	taskService tasks_service.TaskService,
	authMiddleware core_http_middleware.Middleware,
) *TaskHTTPHandler {
	return &TaskHTTPHandler{
		taskService:    taskService,
		authMiddleware: authMiddleware,
	}
}

func (h *TaskHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodPost,
			Path:       "/tasks",
			Handler:    h.CreateTask,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/tasks/{id}",
			Handler:    h.GetTask,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/lists/{id}/tasks",
			Handler:    h.GetListTasks,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodPatch,
			Path:       "/tasks/{id}",
			Handler:    h.PatchTask,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodDelete,
			Path:       "/tasks/{id}",
			Handler:    h.DeleteTask,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodPatch,
			Path:       "/lists/{id}/tasks/reorder",
			Handler:    h.ReorderTasks,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
		{
			Method:     http.MethodPatch,
			Path:       "/tasks/{id}/move",
			Handler:    h.MoveTask,
			Middleware: []core_http_middleware.Middleware{h.authMiddleware},
		},
	}
}
