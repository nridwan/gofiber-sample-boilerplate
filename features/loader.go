package features

import (
	"github.com/nridwan/features/auth"
	"github.com/nridwan/features/guest"

	"github.com/gofiber/fiber/v2"
)

//Register module routes
func Register(app *fiber.App) {
	// app.Use(recover.New())
	auth.Register(app)
	guest.Register(app)
}
