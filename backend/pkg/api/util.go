package api

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func JSON(ctx *fasthttp.RequestCtx, status int, body interface{}) {
	bs, err := json.Marshal(body)
	if err != nil {
		ctx.WriteString("err")
		return
	}

	ctx.Write(bs)
}
