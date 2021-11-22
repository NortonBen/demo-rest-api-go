package bootstrap

import (
	"apm/api"
	"apm/middleware"
	"apm/pkg/migrations"
	server2 "apm/pkg/server"
	"apm/services"
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func Load(db *pg.DB, server *server2.Server, accessSecret string, ctx context.Context) error {

	err := migrations.CreateSchema(db)
	if err != nil {
		return err
	}

	authService := services.NewAuthService(db, accessSecret)
	userService := services.NewUserService(db)

	server.Middleware(
		middleware.NewAuthMiddleware(accessSecret),
	)

	server.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowCredentials: true,
	}))

	// not need  token jwt
	server.Add(
		api.NewAuth(authService),
	)

	// need token jwt
	server.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(accessSecret),
	}))

	server.Add(
		api.NewUser(userService),
	)

	return nil
}
