package contract

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Prefix() string
	Middlewares() []interface{}
	Register(app fiber.Router)
}