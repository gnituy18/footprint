package api

import (
	"encoding/json"
	"net/http"

	"github.com/qiangxue/fasthttp-routing"
)

func JSON(rctx *routing.Context, status int, body interface{}) {
	bs, err := json.Marshal(body)
	if err != nil {
		rctx.Error("json.Marshal failed", http.StatusInternalServerError)
		return
	}

	rctx.Write(bs)
}
