//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tjaszai/go-ms-gateway/internal/db"
	"github.com/tjaszai/go-ms-gateway/internal/http/controller"
	"github.com/tjaszai/go-ms-gateway/internal/http/middleware"
	"github.com/tjaszai/go-ms-gateway/internal/http/server"
	"github.com/tjaszai/go-ms-gateway/internal/repository"
	"github.com/tjaszai/go-ms-gateway/internal/service"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		db.NewDatabaseManager,
		controller.NewDefaultController,
		controller.NewGatewayController,
		controller.NewMicroserviceController,
		controller.NewSecurityController,
		controller.NewUserController,
		middleware.NewAdminGuardMiddleware,
		middleware.NewAuthMiddleware,
		repository.NewMicroserviceRepository,
		repository.NewUserRepository,
		service.NewSecurityService,
		service.NewValidator,
		server.NewServer,
	)
	return &server.Server{}, nil
}
