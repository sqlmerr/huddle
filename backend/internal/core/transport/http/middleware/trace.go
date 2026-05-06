package core_http_middleware

import (
	"net/http"
	"time"

	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	"go.uber.org/zap"
)

func Trace() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)

			before := time.Now()
			log.Debug(">>> incoming HTTP request", zap.Time("time", time.Now().UTC()))

			h.ServeHTTP(w, r)

			log.Debug("<<< done HTTP request", zap.Duration("latency", time.Since(before)))
		})
	}
}
