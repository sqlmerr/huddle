package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.ErrorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) ErrorResponse(status int, err error, msg string) {
	h.w.WriteHeader(status)
	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}
	if err := json.NewEncoder(h.w).Encode(response); err != nil {
		h.log.Error("write HTTP Response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) JSONResponse(status int, data any) {
	h.w.WriteHeader(status)
	if err := json.NewEncoder(h.w).Encode(data); err != nil {
		h.log.Error("write HTTP Response", zap.Error(err))
	}
}
