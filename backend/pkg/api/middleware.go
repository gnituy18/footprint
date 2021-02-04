package api

import (
	"github.com/qiangxue/fasthttp-routing"

	"footprint/pkg/context"
)

func WithContext(ctx *routing.Context) error {
	ctx.SetUserValue("ctx", context.Background())
	ctx.Next()
	return nil
}
