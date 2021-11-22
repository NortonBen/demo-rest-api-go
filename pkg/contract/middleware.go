package contract

import "github.com/gofiber/fiber/v2"

type MiddlewareType string

const (
	MiddlewareLocal MiddlewareType = "local"
	MiddlewareGlobal MiddlewareType = "global"
)

type Middleware interface {
	Name() string
	Type() MiddlewareType
	Handler() fiber.Handler
}
