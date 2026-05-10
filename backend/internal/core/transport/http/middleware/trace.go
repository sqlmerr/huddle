package core_http_middleware

import (
	"net/http"
	"time"

	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
	"go.uber.org/zap"
)

func Trace() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)

			before := time.Now()
			log.Debug(
				">>> incoming HTTP request",
				zap.Time("time", time.Now().UTC()),
				zap.String("http_method", r.Method),
			)

			rw := core_http_response.NewResponseWriter(w)

			h.ServeHTTP(rw, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Duration("latency", time.Since(before)),
				zap.Int("status_code", rw.GetStatusCode()),
			)
		})
	}
}
