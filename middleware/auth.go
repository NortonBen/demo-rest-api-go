package middleware

import (
	"apm/api/common"
	"apm/pkg/contract"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	accessSecret string
}

func NewAuthMiddleware(accessSecret string) *AuthMiddleware {
	return &AuthMiddleware{accessSecret: accessSecret}
}

func (a AuthMiddleware) Name() string {
	return "auth"
}

func (a AuthMiddleware) Type() contract.MiddlewareType {
	return contract.MiddlewareLocal
}

func (a AuthMiddleware) Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)

		if user == nil {
			return c.Status(401).JSON(common.Result{
				Code: 401,
				Message: "not_token",
			})
		}
		return c.Next()
	}
}

