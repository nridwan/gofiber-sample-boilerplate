package features

import (
	"github.com/nridwan/features/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//Register module routes
func Register(app *fiber.App) {
	app.Use(recover.New())
	auth.Register(app)
}
