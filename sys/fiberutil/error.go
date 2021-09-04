package fiberutil

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nridwan/core/data/response"
	"gopkg.in/guregu/null.v3"
)

//AssertError check if connection should be continued or not
func AssertError(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(response.Response{
		Meta: response.Meta{
			Code:    null.IntFrom(int64(code)),
			Message: null.StringFrom(err.Error()),
		},
	})
}
