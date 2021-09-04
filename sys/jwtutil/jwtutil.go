package jwtutil

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nridwan/config/configutil"
	"github.com/nridwan/core/data/response"
	"gopkg.in/guregu/null.v3"
)

var secret = ""
var handler fiber.Handler = nil

func errorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.Status(401).JSON(response.Response{
		Meta: response.Meta{
			Code:    null.IntFrom(401),
			Message: null.StringFrom(err.Error()),
		},
	})
}

func successHandler(ctx *fiber.Ctx) error {
	// claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return ctx.Next()
	// return fiber.NewError(401, "Token is expired")
}

func LoadConfiguration() {
	secret = configutil.Getenv("JWT_SECRET", "")
	asd := jwtware.New(jwtware.Config{
		SigningKey:     []byte(secret),
		ErrorHandler:   errorHandler,
		SuccessHandler: successHandler,
	})
	handler = asd
}

func GetSecret() string {
	return secret
}

func GetHandler() fiber.Handler {
	return handler
}

func GenerateUserToken(id string, apps string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["apps"] = apps
	claims["exp"] = time.Now().Add(time.Minute).Unix()
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(secret))
}
