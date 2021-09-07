package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nridwan/core/middlewares/jwtapp"
	"github.com/nridwan/core/middlewares/jwtuser"
	"github.com/nridwan/sys/jwtutil"
)

const prefix = "auth"

//Register module routes
func Register(app *fiber.App) {
	app.Post(prefix+"/login", jwtutil.GetHandler(), jwtapp.CanAccess, handlerLogin)
	app.Post(prefix+"/refresh", jwtutil.GetHandler(), jwtuser.CanRefresh, handlerRefresh)
	app.Post(prefix+"/logout", jwtutil.GetHandler(), jwtuser.CanAccess, handlerLogout)
	app.Get(prefix+"/profile", jwtutil.GetHandler(), jwtuser.CanAccess, handlerProfile)
}
