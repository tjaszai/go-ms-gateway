//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tjaszai/go-ms-gateway/internal/db"
	"github.com/tjaszai/go-ms-gateway/internal/http/controller"
	"github.com/tjaszai/go-ms-gateway/internal/repository"
	"github.com/tjaszai/go-ms-gateway/internal/server"
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
		repository.NewMicroserviceRepository,
		repository.NewUserRepository,
		service.NewSecurityService,
		service.NewModelValidator,
		server.NewServer,
	)
	return &server.Server{}, nil
}
