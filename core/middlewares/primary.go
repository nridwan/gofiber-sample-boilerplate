package middlewares

import "github.com/gofiber/fiber/v2"

//CheckFalse dummy middleware that always goes wrong
func CheckFalse(ctx *fiber.Ctx) error {
	ctx.Status(404)
	return ctx.SendString("Not Found")
}

//CheckTrue dummy middleware that never goes wrong
func CheckTrue(ctx *fiber.Ctx) error {
	return ctx.Next()
}
