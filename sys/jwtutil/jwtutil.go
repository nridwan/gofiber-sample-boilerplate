package jwtutil

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
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

func LoadConfiguration() {
	secret = configutil.Getenv("JWT_SECRET", "")
	handler = jwtware.New(jwtware.Config{
		SigningKey:   []byte(secret),
		ErrorHandler: errorHandler,
	})
}

func GetSecret() string {
	return secret
}

func GetHandler() fiber.Handler {
	return handler
}

func GetInt64Claim(data interface{}) int64 {
	switch exp := data.(type) {
	case float64:
		return int64(exp)
	case json.Number:
		v, _ := exp.Int64()
		return v
	}
	return 0
}

func GetUint64Claim(data interface{}) uint64 {
	switch exp := data.(type) {
	case float64:
		return uint64(exp)
	case json.Number:
		v, _ := exp.Int64()
		return uint64(v)
	}
	return 0
}
