package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/tjaszai/go-ms-gateway/config"
	_ "github.com/tjaszai/go-ms-gateway/docs"
	"github.com/tjaszai/go-ms-gateway/internal/http/controller"
	"log"
)

type Server struct {
	App                    *fiber.App
	DefaultController      *controller.DefaultController
	MicroserviceController *controller.MicroserviceController
}

func NewServer(dc *controller.DefaultController, mc *controller.MicroserviceController) *Server {
	server := &Server{
		App:                    fiber.New(),
		DefaultController:      dc,
		MicroserviceController: mc,
	}
	server.setupRoutes()
	server.setupMiddlewares()

	return server
}

func (s *Server) setupRoutes() {
	s.setupDefaultRoutes()

	apiRouter := s.App.Group("/api")
	s.setupMSRoutes(apiRouter)
}

func (s *Server) setupDefaultRoutes() {
	s.App.Get("/", s.DefaultController.Index)
	s.App.Get("/docs/*", swagger.HandlerDefault)
	s.App.Get("/HealthCheck", s.DefaultController.HealthCheck)
}

func (s *Server) setupMSRoutes(router fiber.Router) {
	msRouter := router.Group("/Microservices")
	msRouter.Post("/", s.MicroserviceController.Create)
	msRouter.Get("/:id", s.MicroserviceController.GetOne)
	msRouter.Put("/:id", s.MicroserviceController.Update)
	msRouter.Delete("/:id", s.MicroserviceController.Delete)
	msRouter.Get("/", s.MicroserviceController.GetAll)
}

func (s *Server) setupMiddlewares() {
	s.App.Use(logger.New())
}

func (s *Server) Run() {
	port := config.Config("APP_PORT", "3000")
	log.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(s.App.Listen(":" + port))
}
