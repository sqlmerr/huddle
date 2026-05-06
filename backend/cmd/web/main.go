package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sqlmerr/huddle/backend/internal/core/logger"
	core_http_middleware "github.com/sqlmerr/huddle/backend/internal/core/transport/http/middleware"
	core_http_server "github.com/sqlmerr/huddle/backend/internal/core/transport/http/server"
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

	userTransportHTTP := users_transport_http.NewUsersHTTPHandler(nil)
	usersRoutes := userTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersionV1)
	apiVersionRouter.AddRoutes(usersRoutes...)

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
