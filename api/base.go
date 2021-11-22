package api

import (
	"apm/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUser(c *fiber.Ctx) *entities.User {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	firstName := claims["first_name"].(string)
	lastName := claims["last_name"].(string)
	id := int64(claims["id"].(float64))
	tenantId := int64(claims["tenant_id"].(float64))
	departmentId := int64(claims["department_id"].(float64))

	return &entities.User{
		Id:           id,
		Username:     username,
		FirstName:    firstName,
		LastName:     lastName,
		TenantId:     tenantId,
		DepartmentId: departmentId,
	}
}
