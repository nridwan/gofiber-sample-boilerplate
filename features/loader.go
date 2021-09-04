package features

import (
	"github.com/nridwan/features/auth"

	"github.com/gofiber/fiber/v2"
)

//Register module routes
func Register(app *fiber.App) {
	auth.Register(app)
}
