package jwtapp

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/nridwan/core/data/jwtmodel"
	"github.com/nridwan/core/middlewares/jwtuser"
	"github.com/nridwan/models"
	"github.com/nridwan/sys/hashutil"
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

	if (refresh && claims["prv"] != prvRefresh) || (!refresh && claims["prv"] != prv) {
		return false
	}
	var id uint64 = 0
	if res, err := hashutil.DecodeSingle(claims["sub"].(string)); err != nil {
		return false
	} else {
		id = uint64(res)
	}
	exist, err := models.AppTokens(
		qm.Where("app_id=?", id),
		qm.And("hash=?", claims["jti"])).ExistsG(c.Context())
	if err != nil || !exist {
		return false
	}
	if !refresh {
		c.Locals("appId", id)
	}
	return true
}

func CanAccess(c *fiber.Ctx) error {
	if checkUser(c, false) {
		return c.Next()
	}
	return fiber.NewError(401, "Missing or malformed JWT")
}

func MaybeAccess(c *fiber.Ctx) error {
	if jwtuser.CheckUser(c, false) {
		return c.Next()
	}
	return CanAccess(c)
}

func CanRefresh(c *fiber.Ctx) error {
	if checkUser(c, true) {
		return c.Next()
	}
	return fiber.NewError(401, "Missing or malformed JWT")
}

func Logout(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	var uid uint64 = 0
	if res, err := hashutil.DecodeSingle(claims["sub"].(string)); err != nil {
		return err
	} else {
		uid = uint64(res)
	}
	data, err := models.AppTokens(
		qm.Where("app_id=?", uid),
		qm.And("hash=?", claims["jti"])).OneG(ctx.Context())
	if err != nil {
		return err
	}
	data.DeleteG(ctx.Context())
	return nil
}

func Refresh(ctx *fiber.Ctx) (*jwtmodel.TokenResponse, error) {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	var uid uint64 = 0
	if res, err := hashutil.DecodeSingle(claims["sub"].(string)); err != nil {
		return nil, err
	} else {
		uid = uint64(res)
	}
	data, err := models.AppTokens(
		qm.Where("app_id=?", uid),
		qm.And("hash=?", claims["jti"])).OneG(ctx.Context())
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	// Generate encoded token and send it as response.
	accessToken, err := generateAccessTokenHashed(ctx.Context(), claims["sub"].(string), id, now)
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateRefreshTokenHashed(ctx.Context(), claims["sub"].(string), id, now)
	if err != nil {
		return nil, err
	}
	data.Hash = id
	data.ExpiredAt = null.TimeFrom(time.Unix(now, 0).Add(time.Minute * jwtutil.RefreshLifetime))
	data.UpdateG(ctx.Context(), boil.Infer())
	return &jwtmodel.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func GenerateToken(ctx context.Context, sub uint64) (*jwtmodel.TokenResponse, error) {
	now := time.Now().Unix()
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	// Generate encoded token and send it as response.
	accessToken, err := generateAccessToken(ctx, sub, id, now)
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateRefreshToken(ctx, sub, id, now)
	if err != nil {
		return nil, err
	}
	var saved = models.AppToken{
		AppID:     null.Uint64From(sub),
		Hash:      id,
		ExpiredAt: null.TimeFrom(time.Unix(now, 0).Add(time.Minute * jwtutil.RefreshLifetime)),
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

func generateAccessToken(ctx context.Context, sub uint64, jti string, now int64) (string, error) {
	if res, err := hashutil.EncodeSingle(int64(sub)); err != nil {
		return "", err
	} else {
		return generateAccessTokenHashed(ctx, res, jti, now)
	}
}

func generateAccessTokenHashed(ctx context.Context, sub string, jti string, now int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["jti"] = jti
	claims["iat"] = now
	claims["nbf"] = now
	claims["exp"] = time.Unix(now, 0).Add(time.Minute * jwtutil.Lifetime).Unix()
	claims["prv"] = prv
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(jwtutil.GetSecret()))
}

func generateRefreshToken(ctx context.Context, sub uint64, jti string, now int64) (string, error) {
	if res, err := hashutil.EncodeSingle(int64(sub)); err != nil {
		return "", err
	} else {
		return generateRefreshTokenHashed(ctx, res, jti, now)
	}
}

func generateRefreshTokenHashed(ctx context.Context, sub string, jti string, now int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	currentTime := time.Now()
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["jti"] = jti
	claims["iat"] = currentTime.Unix()
	claims["nbf"] = currentTime.Unix()
	claims["exp"] = time.Unix(now, 0).Add(time.Minute * jwtutil.RefreshLifetime).Unix()
	claims["prv"] = prvRefresh
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(jwtutil.GetSecret()))
}
