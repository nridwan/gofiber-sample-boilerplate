package features

import (
	"github.com/nridwan/config/configutil"
	"github.com/nridwan/features/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//Register module routes
func Register(app *fiber.App) {
	app.Use(recover.New())
	corsOrigin := configutil.Getenv("CORS_ORIGIN", "")
	if corsOrigin != "" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: corsOrigin,
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	}
	auth.Register(app)
}
