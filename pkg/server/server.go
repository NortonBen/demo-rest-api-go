package server

import (
	_ "apm/docs"
	"apm/pkg/contract"
	"context"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"log"
	"net"
	"os"
	"strings"
)

type Server struct {
	app           *fiber.App
	mapMiddleware map[string]fiber.Handler
	ctx           context.Context
}

func NewServer(ctx context.Context) *Server {
	app := fiber.New()
	env, _ := os.LookupEnv("ENV_GO")
	if strings.ToLower(env) != "product" {
		app.Get("/swagger/*", swagger.Handler)
	}

	return &Server{
		app: app, ctx: ctx,
		mapMiddleware: map[string]fiber.Handler{},
	}
}

func (s *Server) Listener(ln net.Listener) error {
	return s.app.Listener(ln)
}

func (s *Server) Add(controllers ...contract.Controller) {
	for _, controller := range controllers {
		var middlewares []fiber.Handler
		if controller.Middlewares() != nil {
			middlewares = s.getMiddleware(controller.Middlewares())
		}
		prefix := controller.Prefix()
		group := s.app.Group(prefix, middlewares...)
		controller.Register(group)
	}
}

func (s *Server) getMiddleware(list []interface{}) []fiber.Handler {
	middlewares := make([]fiber.Handler, 0)
	for _, item := range list {
		switch item.(type) {
		case string:
			{
				middleware, success := s.mapMiddleware[item.(string)]
				if success {
					middlewares = append(middlewares, middleware)
				}
			}
		case fiber.Handler:
			{
				middlewares = append(middlewares, item.(fiber.Handler))
			}
		default:
			{
				log.Fatalf("not found type middleware add %v", item)
			}
		}
	}

	return middlewares
}

func (s *Server) Middleware(middlewares ...contract.Middleware) {
	for _, middleware := range middlewares {
		if middleware.Type() == contract.MiddlewareGlobal {
			s.app.Use(middleware.Handler())
		} else {
			s.mapMiddleware[middleware.Name()] = middleware.Handler()
		}
	}
}

func (s *Server) Use(fn fiber.Handler) {
	s.app.Use(fn)
}

func (s *Server) Run(address string) error {
	return s.app.Listen(address)
}
