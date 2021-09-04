package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nridwan/models"
	"github.com/nridwan/sys/dbutil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func handlerLogin(ctx *fiber.Ctx) error {
	data, err := models.Users(qm.Where("username=?", ctx.Query("username"))).One(ctx.Context(), dbutil.Default())
	success := err == nil
	success = success && bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(ctx.Query("test"))) == nil
	if !success {
		data = nil
	}

	return ctx.JSON(map[string]interface{}{
		"success": success,
		"data":    data,
	})
}
