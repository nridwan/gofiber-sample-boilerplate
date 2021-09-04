package handler

import "github.com/valyala/fasthttp"

//AssertError check if connection should be continued or not
func AssertError(ctx *fasthttp.RequestCtx, err error) bool {
	if err != nil {
		ctx.Error(err.Error(), 500)
		return true
	}
	return false
}
