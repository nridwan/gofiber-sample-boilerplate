package auth

import (
	"github.com/gofiber/fiber/v2"
)

const prefix = "auth"

//Register module routes
func Register(app *fiber.App) {
	app.Post(prefix+"/login", handlerLogin)
}
