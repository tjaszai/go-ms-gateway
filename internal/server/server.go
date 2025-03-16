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
	GatewayController      *controller.GatewayController
	MicroserviceController *controller.MicroserviceController
	SecurityController     *controller.SecurityController
	UserController         *controller.UserController
}

func NewServer(dc *controller.DefaultController,
	gc *controller.GatewayController,
	mc *controller.MicroserviceController,
	sc *controller.SecurityController,
	uc *controller.UserController) *Server {
	server := &Server{
		App:                    fiber.New(),
		DefaultController:      dc,
		GatewayController:      gc,
		MicroserviceController: mc,
		SecurityController:     sc,
		UserController:         uc,
	}
	server.setupRoutes()
	server.setupMiddlewares()

	return server
}

func (s *Server) setupRoutes() {
	apiRouter := s.App.Group("/api")
	s.setupDefaultRoutes()
	s.setupGatewayRoutes(apiRouter)
	s.setupMSRoutes(apiRouter)
	s.setupSecurityRoutes(apiRouter)
	s.setupUserRoutes(apiRouter)
}

func (s *Server) setupDefaultRoutes() {
	s.App.Get("/", s.DefaultController.Index)
}

func (s *Server) setupGatewayRoutes(router fiber.Router) {
	router.Get("/docs/*", swagger.HandlerDefault)
	router.Get("/HealthCheck", s.GatewayController.HealthCheck)
	router.Post("/CallMs", s.GatewayController.CallMs)
}

func (s *Server) setupMSRoutes(router fiber.Router) {
	msRouter := router.Group("/Microservices")
	msRouter.Post("/", s.MicroserviceController.Create)
	msRouter.Get("/:id", s.MicroserviceController.GetOne)
	msRouter.Put("/:id", s.MicroserviceController.Update)
	msRouter.Delete("/:id", s.MicroserviceController.Delete)
	msRouter.Get("/", s.MicroserviceController.GetAll)
}

func (s *Server) setupSecurityRoutes(router fiber.Router) {
	msRouter := router.Group("/Auth")
	msRouter.Post("/Login", s.SecurityController.Login)
}

func (s *Server) setupUserRoutes(router fiber.Router) {
	msRouter := router.Group("/Users")
	msRouter.Post("/", s.UserController.Create)
	msRouter.Get("/:id", s.UserController.GetOne)
	msRouter.Put("/:id", s.UserController.Update)
	msRouter.Delete("/:id", s.UserController.Delete)
	msRouter.Get("/", s.UserController.GetAll)
}

func (s *Server) setupMiddlewares() {
	s.App.Use(logger.New())
}

func (s *Server) Run() {
	port := config.Config("APP_PORT", "3000")
	log.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(s.App.Listen(":" + port))
}
