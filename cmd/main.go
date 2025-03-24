package main

import "log"

// @title                       Go Microservice Gateway
// @version                     0.1.0
// @description                 The MSG provides tools for managing microservices.
// @BasePath                    /
// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
func main() {
	server, err := InitializeServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}
	server.Run()
}
