package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	auth_postgres_repository "github.com/sqlmerr/huddle/backend/internal/auth/repository/postgres"
	auth_service "github.com/sqlmerr/huddle/backend/internal/auth/service"
	auth_transport_http "github.com/sqlmerr/huddle/backend/internal/auth/transport/http"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_postgres_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
	spaces_postgres_repository "github.com/sqlmerr/huddle/backend/internal/spaces/repository/postgres"
	spaces_service "github.com/sqlmerr/huddle/backend/internal/spaces/service"
	spaces_http_transport "github.com/sqlmerr/huddle/backend/internal/spaces/transport/http"
	users_postgres_repository "github.com/sqlmerr/huddle/backend/internal/users/repository/postgres"
	users_service "github.com/sqlmerr/huddle/backend/internal/users/service"
	users_transport_http "github.com/sqlmerr/huddle/backend/internal/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log, err := logger.New(logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init logger:", err)
		os.Exit(1)
	}
	defer log.Close()

	log.Debug("Starting Huddle backend!")

	postgresConfig := core_postgres_pool.LoadConfigMust()
	pool, err := core_postgres_pool.NewPool(ctx, *postgresConfig)
	if err != nil {
		log.Error("failed to create postgres pool", zap.Error(err))
		os.Exit(1)
	}
	authConfig := core_auth.LoadConfigMust()
	jwtProcessor := core_auth.NewJWTProcessor(*authConfig)
	authMiddleware := core_http_middleware.Auth(*jwtProcessor)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersionV1)

	log.Debug("feature initialization", zap.String("feature", "users"))
	userRepository := users_postgres_repository.NewUserRepository(pool)
	userService := users_service.NewUserService(userRepository)
	userTransportHTTP := users_transport_http.NewUserHTTPHandler(userService, authMiddleware)
	apiVersionRouter.AddRoutes(userTransportHTTP.Routes()...)

	log.Debug("feature initialization", zap.String("feature", "auth"))
	authRepository := auth_postgres_repository.NewAuthRepository(pool)
	authService := auth_service.NewAuthService(authRepository, jwtProcessor)
	authTransportHTTP := auth_transport_http.NewAuthHTTPHandler(authService, authMiddleware)
	apiVersionRouter.AddRoutes(authTransportHTTP.Routes()...)

	log.Debug("feature initialization", zap.String("feature", "spaces"))
	spaceRepository := spaces_postgres_repository.NewSpaceRepository(pool)
	spaceService := spaces_service.NewSpaceService(spaceRepository)
	spaceTransportHTTP := spaces_http_transport.NewSpaceHTTPHandler(spaceService, authMiddleware)
	apiVersionRouter.AddRoutes(spaceTransportHTTP.Routes()...)

	httpServer := core_http_server.NewHttpServer(
		*core_http_server.LoadConfigMust(),
		log,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(log),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	httpServer.RegisterRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("failed to run http server", zap.Error(err))
		os.Exit(1)
	}
}
