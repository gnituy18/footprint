package api

import (
	"net/http"

	"encoding/json"

	"github.com/valyala/fasthttp"
)

func JSON(ctx *fasthttp.RequestCtx, status int, body interface{}) {
	bs, err := json.Marshal(body)
	if err != nil {
		ctx.Error("json.Marshal failed", http.StatusInternalServerError)
		return
	}

	ctx.Write(bs)
}
