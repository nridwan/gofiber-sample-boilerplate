package guest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nridwan/core/middlewares/jwtapp"
	"github.com/nridwan/sys/jwtutil"
)

const prefix = "guests"

//Register module routes
func Register(app *fiber.App) {
	app.Post(prefix+"/login", handlerLogin)
	app.Post(prefix+"/refresh", jwtutil.GetHandler(), jwtapp.CanRefresh, handlerRefresh)
	app.Post(prefix+"/logout", jwtutil.GetHandler(), jwtapp.CanAccess, handlerLogout)
	app.Get(prefix+"/profile", jwtutil.GetHandler(), jwtapp.CanAccess, handlerProfile)
}
