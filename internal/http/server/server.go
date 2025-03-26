package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/tjaszai/go-ms-gateway/config"
	_ "github.com/tjaszai/go-ms-gateway/docs"
	"github.com/tjaszai/go-ms-gateway/internal/http/controller"
	"github.com/tjaszai/go-ms-gateway/internal/http/middleware"
	"log"
)

type Server struct {
	App                           *fiber.App
	DefaultController             *controller.DefaultController
	GatewayController             *controller.GatewayController
	MicroserviceController        *controller.MicroserviceController
	MicroserviceVersionController *controller.MicroserviceVersionController
	SecurityController            *controller.SecurityController
	UserController                *controller.UserController
	AdminGuardMiddleware          *middleware.AdminGuardMiddleware
	AuthMiddleware                *middleware.AuthMiddleware
}

func NewServer(
	dc *controller.DefaultController,
	gc *controller.GatewayController,
	mc *controller.MicroserviceController,
	mvc *controller.MicroserviceVersionController,
	sc *controller.SecurityController,
	uc *controller.UserController,
	agm *middleware.AdminGuardMiddleware,
	am *middleware.AuthMiddleware) *Server {
	server := &Server{
		App:                           fiber.New(),
		DefaultController:             dc,
		GatewayController:             gc,
		MicroserviceController:        mc,
		MicroserviceVersionController: mvc,
		SecurityController:            sc,
		UserController:                uc,
		AdminGuardMiddleware:          agm,
		AuthMiddleware:                am,
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
	router.Post("/CallMs", s.AuthMiddleware.Check, s.GatewayController.CallMs)
	router.Get("/docs/*", swagger.HandlerDefault)
	router.Get("/HealthCheck", s.GatewayController.HealthCheck)
}

func (s *Server) setupMSRoutes(router fiber.Router) {
	msRouter := router.Group("/Microservices", s.AuthMiddleware.Check)
	msRouter.Get("/", s.MicroserviceController.GetAll)
	msRouter.Post("/", s.AdminGuardMiddleware.Check, s.MicroserviceController.Create)
	msRouter.Get("/:id", s.MicroserviceController.GetOne)
	msRouter.Put("/:id", s.AdminGuardMiddleware.Check, s.MicroserviceController.Update)
	msRouter.Delete("/:id", s.AdminGuardMiddleware.Check, s.MicroserviceController.Delete)

	msVersionRouter := msRouter.Group("/:id/Versions", s.AuthMiddleware.Check)
	msVersionRouter.Get("/", s.MicroserviceVersionController.GetAll)
	msVersionRouter.Post("/", s.AdminGuardMiddleware.Check, s.MicroserviceVersionController.Create)
	msVersionRouter.Get("/:vid", s.MicroserviceVersionController.GetOne)
	msVersionRouter.Put("/:vid", s.AdminGuardMiddleware.Check, s.MicroserviceVersionController.Update)
	msVersionRouter.Delete("/:vid", s.AdminGuardMiddleware.Check, s.MicroserviceVersionController.Delete)
}

func (s *Server) setupSecurityRoutes(router fiber.Router) {
	secRouter := router.Group("/Security")
	secRouter.Post("/Login", s.SecurityController.Login)
}

func (s *Server) setupUserRoutes(router fiber.Router) {
	userRouter := router.Group("/Users", s.AuthMiddleware.Check)
	userRouter.Get("/", s.UserController.GetAll)
	userRouter.Post("/", s.AdminGuardMiddleware.Check, s.UserController.Create)
	userRouter.Get("/:id", s.UserController.GetOne)
	userRouter.Put("/:id", s.AdminGuardMiddleware.Check, s.UserController.Update)
	userRouter.Delete("/:id", s.AdminGuardMiddleware.Check, s.UserController.Delete)
}

func (s *Server) setupMiddlewares() {
	s.App.Use(logger.New())
}

func (s *Server) Run() {
	port := config.Config("APP_PORT", "3000")
	log.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(s.App.Listen(":" + port))
}
