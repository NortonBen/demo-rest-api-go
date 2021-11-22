package api

import (
	"apm/api/common"
	"apm/model"
	"apm/services"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	authService *services.AuthService
}

func NewAuth(authService *services.AuthService) *Auth {
	return &Auth{authService: authService}
}

func (a Auth) Middlewares() []interface{} {
	return nil
}

func (a Auth) Prefix() string {
	return "auth"
}

func (a *Auth) Register(app fiber.Router) {
	app.Post("login", a.Login)
}

// Login
// @Summary Login system
// @Description Login with username and password
// @Tags login
// @Accept  json
// @Produce  json
// @Param   message body model.Login true  "Login"
// @Success 200 {object} common.Result{document=model.LoginResult} "ok"
// @Failure default {object} common.Result "Error"
// @Router /auth/login [post]
func (a Auth) Login(c *fiber.Ctx) error {
	login := new(model.Login)
	if err := c.BodyParser(login); err != nil {
		return err
	}

	errors := model.ValidateStruct(*login)
	if errors != nil {
		return common.ResultValidate(c, errors)
	}

	rs, err := a.authService.Login(login)
	if err != nil {
		return common.ResultError(c, err)
	}

	return common.ResultSuccess(c, rs)
}
