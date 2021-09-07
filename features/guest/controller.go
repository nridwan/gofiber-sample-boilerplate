package guest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nridwan/core/data/response"
	"github.com/nridwan/core/middlewares/jwtapp"
	"github.com/nridwan/models"
	"github.com/nridwan/sys/dbutil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func handlerLogin(ctx *fiber.Ctx) error {
	request := paramApps{}
	ctx.BodyParser(&request)
	if len(request.Alias.String) == 0 {
		return ctx.JSON(response.CreateMetaResponse(500, "failed", []response.Error{{
			Code:   "alias",
			Reason: "Alias must not be empty",
		}}))
	}
	if len(request.Appkey.String) == 0 {
		return ctx.JSON(response.CreateMetaResponse(500, "failed", []response.Error{{
			Code:   "appkey",
			Reason: "Appkey must not be empty",
		}}))
	}
	data, err := models.APIApps(
		qm.Where("alias=?", request.Alias.String),
		qm.And("appkey=?", request.Appkey.String)).One(ctx.Context(), dbutil.Default())
	if err != nil {
		return ctx.JSON(response.CreateMetaResponse(500, "Wrong Username / Password", []response.Error{{
			Code:   "not_found",
			Reason: "Wrong Alias / Appkey",
		}}))
	}

	t, err := jwtapp.GenerateToken(ctx.Context(), data.ID)
	if err != nil {
		return ctx.JSON(response.CreateMetaResponse(500, "", []response.Error{}))
	}

	return ctx.JSON(response.CreateResponse(200, "success", t))
}

func handlerRefresh(ctx *fiber.Ctx) error {
	t, err := jwtapp.Refresh(ctx)
	if err != nil {
		return ctx.JSON(response.CreateMetaResponse(500, "", []response.Error{}))
	}

	return ctx.JSON(response.CreateResponse(200, "success", t))
}

func handlerLogout(ctx *fiber.Ctx) error {
	err := jwtapp.Logout(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(response.CreateResponse(200, "success", nil))
}

func handlerProfile(ctx *fiber.Ctx) error {
	user := ctx.Locals("userData")
	return ctx.JSON(response.CreateResponse(200, "success", user))
}
