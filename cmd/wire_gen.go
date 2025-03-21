// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/tjaszai/go-ms-gateway/internal/db"
	"github.com/tjaszai/go-ms-gateway/internal/http/controller"
	"github.com/tjaszai/go-ms-gateway/internal/repository"
	"github.com/tjaszai/go-ms-gateway/internal/server"
	"github.com/tjaszai/go-ms-gateway/internal/service"
)

// Injectors from wire.go:

func InitializeServer() (*server.Server, error) {
	defaultController := controller.NewDefaultController()
	databaseManager := db.NewDatabaseManager()
	gatewayController := controller.NewGatewayController(databaseManager)
	microserviceRepository := repository.NewMicroserviceRepository(databaseManager)
	modelValidator := service.NewModelValidator()
	microserviceController := controller.NewMicroserviceController(microserviceRepository, modelValidator)
	userRepository := repository.NewUserRepository(databaseManager)
	securityService := service.NewSecurityService()
	securityController := controller.NewSecurityController(userRepository, modelValidator, securityService)
	userController := controller.NewUserController(userRepository, modelValidator)
	serverServer := server.NewServer(defaultController, gatewayController, microserviceController, securityController, userController)
	return serverServer, nil
}
