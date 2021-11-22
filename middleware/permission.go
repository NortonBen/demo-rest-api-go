package middleware

import (
	"apm/api/common"
	"apm/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Permissions(list ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		raws := claims["permissions"].([]interface{})

		var permissions = make([]string, 0)

		for _, raw := range raws {
			permissions = append(permissions, raw.(string))
		}

		if util.StringInSlice("own_system", permissions) {
			return c.Next()
		}

		for _, permission := range list {
			if util.StringInSlice(permission, permissions) {
				return c.Next()
			}
		}

		return c.Status(403).JSON(common.Result{
			Code:    403,
			Message: "not_permission",
		})

	}
}
