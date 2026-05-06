package core_http_middleware

import (
	"context"
	"net/http"

	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	"go.uber.org/zap"
)

func Logger(log *logger.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-ID")

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), logger.LoggerKey, l)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
