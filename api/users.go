package api

import (
	"apm/api/common"
	"apm/middleware"
	"apm/model"
	"apm/pkg/entities"
	"apm/pkg/util"
	"apm/services"
	_ "github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	userService *services.UserService
}

func NewUser(userService *services.UserService) *User {
	return &User{userService: userService}
}

func (u User) Prefix() string {
	return "users"
}

func (u User) Middlewares() []interface{} {
	return []interface{}{
		"auth",
	}
}

func (u User) Register(app fiber.Router) {
	app.Get("", u.List).Use(middleware.Permissions("read_user"))
	app.Post("", u.Store).Use(middleware.Permissions("write_user"))
	app.Get("/permission", u.Permission).Use(middleware.Permissions("read_user"))
	app.Get(":id", u.Get).Use(middleware.Permissions("read_user"))
	app.Put(":id", u.Update).Use(middleware.Permissions("write_user"))
	app.Delete(":id", u.Delete).Use(middleware.Permissions("write_user"))
}

// Get List User
// @Summary List User
// @Description List User
// @Tags user
// @Accept  json
// @Produce  json
// @Param search query string false "search"
// @Param limit query number false "limit"  default(50)
// @Param skip query number false "skip" default(0)
// @Success 200 {object} common.Result{document=model.Records} "ok"
// @Failure default {object} common.Result "Error"
// @Security ApiKeyAuth
// @Router /users [get]
func (u User) List(c *fiber.Ctx) error {

	search := c.Query("search")
	skip, limit := common.GetSkipLimit(c)

	rs, err := u.userService.List(&model.SearchBase{
		Limit:  limit,
		Skip:   skip,
		Search: search,
	})
	if err != nil {
		return common.ResultError(c, err)
	}

	return common.ResultSuccess(c, rs)
}

// Get user detail
// @Summary Get user detail
// @Description Get user detail
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id     path    number     true        "User ID"
// @Success 200 {object} common.Result{document=model.UserUpdate} "ok"
// @Failure default {object} common.Result "Error"
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (u User) Get(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	rs, err := u.userService.Get(int64(id))
	if err != nil {
		return common.ResultError(c, err)
	}

	return common.ResultSuccess(c, rs)
}

// User Store
// @Summary User Store
// @Description create user
// @Tags user
// @Accept  json
// @Produce  json
// @Param   message body model.UserCreate true  "Login"
// @Success 200 {object} common.Result "ok"
// @Failure default {object} common.Result "Error"
// @Security ApiKeyAuth
// @Router /users [post]
func (u User) Store(c *fiber.Ctx) error {

	user := GetUser(c)

	data := new(model.UserCreate)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	errors := model.ValidateStruct(*data)
	if errors != nil {
		return common.ResultValidate(c, errors)
	}

	request := new(entities.User)
	util.Mapper(data, request)
	request.TenantId = user.TenantId

	request.Permissions = model.FilterPermission(request.Permissions)

	rs, err := u.userService.Store(request)
	if err != nil {
		return common.ResultError(c, err)
	}

	return common.ResultSuccess(c, rs)
}

// User Update
// @Summary User Update
// @Description update user
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id     path    number     true        "User ID"
// @Param   message body model.UserCreate true  "Login"
// @Success 200 {object} common.Result "ok"
// @Failure default {object} common.Result "Error"
// @Security ApiKeyAuth
// @Router /users/{id} [put]
func (u User) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	data := new(model.UserUpdate)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	errors := model.ValidateStruct(*data)
	if errors != nil {
		return common.ResultValidate(c, errors)
	}

	request := new(entities.User)
	util.Mapper(data, request)

	request.Permissions = model.FilterPermission(request.Permissions)

	rs, err := u.userService.Update(int64(id), request)
	if err != nil {
		return common.ResultError(c, err)
	}

	return common.ResultSuccess(c, rs)
}

// Users Delete
// @Summary Users Delete
// @Description delete Users
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id     path    number     true        "User ID"
// @Success 200 {object} common.Result "ok"
// @Failure default {object} common.Result "Error"
// @Security ApiKeyAuth
// @Router /users/{id} [delete]
func (u User) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	rs, err := u.userService.Delete(int64(id))
	if err != nil {
		return common.ResultError(c, err)
	}

	return common.ResultSuccess(c, rs)
}

// Users Permission
// @Summary Users Permission
// @Description Users Permission
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Result{document=[]model.Permission} "ok"
// @Failure default {object} common.Result "Error"
// @Security ApiKeyAuth
// @Router /users/permission [get]
func (u User) Permission(c *fiber.Ctx) error {
	return common.ResultSuccess(c, model.ListPermission)
}
