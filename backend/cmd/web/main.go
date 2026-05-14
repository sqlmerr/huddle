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
	boards_postgres_repository "github.com/sqlmerr/huddle/backend/internal/boards/repository/postgres"
	boards_service "github.com/sqlmerr/huddle/backend/internal/boards/service"
	boards_http_transport "github.com/sqlmerr/huddle/backend/internal/boards/transport/http"
	core_access "github.com/sqlmerr/huddle/backend/internal/core/access"
	core_auth "github.com/sqlmerr/huddle/backend/internal/core/auth"
	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_pgx_pool "github.com/sqlmerr/huddle/backend/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
	lists_postgres_repository "github.com/sqlmerr/huddle/backend/internal/lists/repository/postgres"
	lists_service "github.com/sqlmerr/huddle/backend/internal/lists/service"
	lists_http_transport "github.com/sqlmerr/huddle/backend/internal/lists/transport/http"
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

	postgresConfig := core_pgx_pool.LoadConfigMust()
	pool, err := core_pgx_pool.NewPool(ctx, *postgresConfig)
	if err != nil {
		log.Error("failed to create postgres pool", zap.Error(err))
		os.Exit(1)
	}
	authConfig := core_auth.LoadConfigMust()
	jwtProcessor := core_auth.NewJWTProcessor(*authConfig)
	authMiddleware := core_http_middleware.Auth(*jwtProcessor)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersionV1)

	userRepository := users_postgres_repository.NewUserRepository(pool)
	authRepository := auth_postgres_repository.NewAuthRepository(pool)
	spaceRepository := spaces_postgres_repository.NewSpaceRepository(pool)
	boardRepository := boards_postgres_repository.NewBoardRepository(pool)
	listRepository := lists_postgres_repository.NewListRepository(pool)
	accessService := core_access.NewAccessService(spaceRepository, boardRepository, listRepository)

	log.Debug("feature initialization", zap.String("feature", "users"))
	userService := users_service.NewUserService(userRepository)
	userTransportHTTP := users_transport_http.NewUserHTTPHandler(userService, authMiddleware)
	apiVersionRouter.AddRoutes(userTransportHTTP.Routes()...)

	log.Debug("feature initialization", zap.String("feature", "auth"))
	authService := auth_service.NewAuthService(authRepository, jwtProcessor)
	authTransportHTTP := auth_transport_http.NewAuthHTTPHandler(authService, authMiddleware)
	apiVersionRouter.AddRoutes(authTransportHTTP.Routes()...)

	log.Debug("feature initialization", zap.String("feature", "spaces"))
	spaceService := spaces_service.NewSpaceService(spaceRepository, accessService)
	spaceTransportHTTP := spaces_http_transport.NewSpaceHTTPHandler(spaceService, authMiddleware)
	apiVersionRouter.AddRoutes(spaceTransportHTTP.Routes()...)

	log.Debug("feature initialization", zap.String("feature", "boards"))
	boardService := boards_service.NewBoardService(boardRepository, accessService)
	boardTransportHTTP := boards_http_transport.NewBoardsHTTPHandler(boardService, authMiddleware)
	apiVersionRouter.AddRoutes(boardTransportHTTP.Routes()...)

	log.Debug("feature initialization", zap.String("feature", "lists"))
	listService := lists_service.NewListService(listRepository, accessService)
	listTransportHTTP := lists_http_transport.NewListHTTPHandler(listService, authMiddleware)
	apiVersionRouter.AddRoutes(listTransportHTTP.Routes()...)

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
