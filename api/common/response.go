package common

import (
	"apm/model"
	"apm/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type Result struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Document interface{} `json:"document"`
}

func ResultError(c *fiber.Ctx, err error) error {
	switch err.(type) {
	case *errors.Error:
		{
			e := err.(*errors.Error)
			return c.Status(e.Code).JSON(Result{
				Code:     e.Code,
				Message:  e.Message,
				Document: e,
			})
		}

	}
	return c.Status(500).JSON(Result{
		Code:     500,
		Message:  err.Error(),
		Document: nil,
	})
}


func ResultValidate(c *fiber.Ctx, errors []*model.ErrorResponse) error {
	return c.Status(400).JSON(Result{
		Code:     400,
		Message:  "validated_data",
		Document: errors,
	})
}

func ResultSuccess(c *fiber.Ctx, data interface{}) error {

	return c.JSON(Result{
		Code:     200,
		Message:  "",
		Document: data,
	})
}
