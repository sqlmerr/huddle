package core_http_middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_response "github.com/sqlmerr/huddle/backend/internal/core/transport/http/response"
)

func Auth(jwtProcessor core_auth.JWTProcessor) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			err := fmt.Errorf("invalid jwt token: %w", core_errors.ErrAccessDenied)

			authorizationHeader := r.Header.Get("Authorization")
			headerParts := strings.SplitN(authorizationHeader, " ", 2)
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				responseHandler.ErrorResponse(err, "invalid jwt token. Must be in format: 'Bearer <jwt token>'")
				return
			}

			jwtToken := headerParts[1]
			userID, err := jwtProcessor.ValidateToken(jwtToken)
			if err != nil {
				responseHandler.ErrorResponse(err, "invalid jwt token")
				return
			}
			ctx = context.WithValue(ctx, core_auth.UserIDContextKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
