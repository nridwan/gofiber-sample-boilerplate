package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nridwan/core/data/response"
	"github.com/nridwan/core/middlewares/jwtuser"
	"github.com/nridwan/models"
	"github.com/nridwan/sys/dbutil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func handlerLogin(ctx *fiber.Ctx) error {
	request := paramLogin{}
	ctx.BodyParser(&request)
	if len(request.Username.String) == 0 {
		return ctx.JSON(response.CreateMetaResponse(500, "failed", []response.Error{{
			Code:   "username",
			Reason: "Username must not be empty",
		}}))
	}
	if len(request.Password.String) == 0 {
		return ctx.JSON(response.CreateMetaResponse(500, "failed", []response.Error{{
			Code:   "password",
			Reason: "Password must not be empty",
		}}))
	}
	data, err := models.Users(qm.Where("username=?", request.Username.String)).One(ctx.Context(), dbutil.Default())
	success := err == nil
	success = success && bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(request.Password.String)) == nil
	if !success {
		return ctx.JSON(response.CreateMetaResponse(500, "Wrong Username / Password", []response.Error{{
			Code:   "not_found",
			Reason: "Wrong Username / Password",
		}}))
	}

	t, err := jwtuser.GenerateToken(ctx.Context(), data.ID)
	if err != nil {
		return ctx.JSON(response.CreateMetaResponse(500, "", []response.Error{}))
	}

	return ctx.JSON(response.CreateResponse(200, "success", t))
}

func handlerRefresh(ctx *fiber.Ctx) error {
	t, err := jwtuser.Refresh(ctx)
	if err != nil {
		return ctx.JSON(response.CreateMetaResponse(500, "", []response.Error{}))
	}

	return ctx.JSON(response.CreateResponse(200, "success", t))
}

func handlerLogout(ctx *fiber.Ctx) error {
	err := jwtuser.Logout(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(response.CreateResponse(200, "success", nil))
}

func handlerProfile(ctx *fiber.Ctx) error {
	user := ctx.Locals("userData")
	return ctx.JSON(response.CreateResponse(200, "success", getUserDto(user.(*models.User))))
}
