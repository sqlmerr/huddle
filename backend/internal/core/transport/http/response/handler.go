package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *logger.Logger
	w   http.ResponseWriter
}

func NewHTTPResponseHandler(
	log *logger.Logger, w http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log, w: w,
	}
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	err := fmt.Errorf("unexpected panic: %v: %w", p, core_errors.ErrInternalServerError)

	h.ErrorResponse(err, msg)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)
	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))
	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}
	h.JSONResponse(statusCode, response)
}

func (h *HTTPResponseHandler) JSONResponse(statusCode int, data any) {
	h.w.WriteHeader(statusCode)
	if err := json.NewEncoder(h.w).Encode(data); err != nil {
		h.log.Error("write HTTP Response", zap.Error(err))
	}
}
