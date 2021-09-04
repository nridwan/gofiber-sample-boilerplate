package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nridwan/sys/jwtutil"
)

const prefix = "auth"

//Register module routes
func Register(app *fiber.App) {
	app.Post(prefix+"/login", handlerLogin)
	app.Get(prefix+"/profile", jwtutil.GetHandler(), handlerProfile)
}
