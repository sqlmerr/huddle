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

	userRepository := users_postgres_repository.NewUsersRepository(pool)
	userService := users_service.NewUserService(userRepository)
	userTransportHTTP := users_transport_http.NewUsersHTTPHandler(userService, authMiddleware)
	apiVersionRouter.AddRoutes(userTransportHTTP.Routes()...)

	authRepository := auth_postgres_repository.NewAuthRepository(pool)
	authService := auth_service.NewAuthService(authRepository, jwtProcessor)
	authTransportHTTP := auth_transport_http.NewAuthHTTPHandler(authService)
	apiVersionRouter.AddRoutes(authTransportHTTP.Routes()...)

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
