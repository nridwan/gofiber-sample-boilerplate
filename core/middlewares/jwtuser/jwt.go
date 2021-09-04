package jwtuser

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/nridwan/core/data/jwtmodel"
	"github.com/nridwan/models"
	"github.com/nridwan/sys/jwtutil"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const identifier = "jwtauth/app"

var sha = sha1.New()
var prv = base64.StdEncoding.EncodeToString(sha.Sum([]byte(identifier)))
var prvRefresh = base64.StdEncoding.EncodeToString(sha.Sum([]byte(identifier + "/refresh")))

func checkUser(c *fiber.Ctx, refresh bool) bool {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	println(claims["prv"] != prv)
	println(claims["prv"] != prvRefresh)

	if (refresh && claims["prv"] != prvRefresh) || (!refresh && claims["prv"] != prv) {
		return false
	}
	exist, err := models.UserTokens(
		qm.Where("user_id=?", claims["sub"]),
		qm.And("hash=?", claims["jti"])).ExistsG(c.Context())
	if err != nil || !exist {
		return false
	}
	if !refresh {
		users, err := models.Users(qm.Where("id=?", claims["sub"])).OneG(c.Context())
		if err != nil {
			return false
		}
		c.Locals("userData", users)
	}
	return true
}

func CanAccess(c *fiber.Ctx) error {
	if checkUser(c, false) {
		return c.Next()
	}
	return fiber.NewError(401, "Missing or malformed JWT")
}

func CanRefresh(c *fiber.Ctx) error {
	if checkUser(c, true) {
		return c.Next()
	}
	return fiber.NewError(401, "Missing or malformed JWT")
}

func Logout(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	data, err := models.UserTokens(
		qm.Where("user_id=?", claims["sub"]),
		qm.And("hash=?", claims["jti"])).OneG(ctx.Context())
	if err != nil {
		return err
	}
	data.DeleteG(ctx.Context())
	return nil
}

func GenerateToken(ctx context.Context, sub uint64, api uint64) (*jwtmodel.TokenResponse, error) {
	now := time.Now().Unix()
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	// Generate encoded token and send it as response.
	accessToken, err := generateAccessToken(ctx, sub, api, id, now)
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateRefreshToken(ctx, sub, api, id, now)
	if err != nil {
		return nil, err
	}
	var saved = models.UserToken{
		UserID: null.Uint64From(sub),
		Hash:   id,
	}
	err = saved.InsertG(ctx, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &jwtmodel.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateAccessToken(ctx context.Context, sub uint64, api uint64, jti string, now int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["api"] = api
	claims["jti"] = jti
	claims["iat"] = now
	claims["nbf"] = now
	claims["exp"] = time.Unix(now, 0).Add(time.Minute).Unix()
	claims["prv"] = prv
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(jwtutil.GetSecret()))
}

func generateRefreshToken(ctx context.Context, sub uint64, api uint64, jti string, now int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	currentTime := time.Now()
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["api"] = api
	claims["jti"] = jti
	claims["iat"] = currentTime.Unix()
	claims["nbf"] = currentTime.Unix()
	claims["exp"] = time.Unix(now, 0).Add(time.Minute).Unix()
	claims["prv"] = prvRefresh
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(jwtutil.GetSecret()))
}
