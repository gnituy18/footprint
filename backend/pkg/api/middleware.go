package api

import (
	"github.com/qiangxue/fasthttp-routing"
	"go.uber.org/zap"

	"footprint/pkg/context"
)

func InjectContext(rctx *routing.Context) error {
	rctx.Set("ctx", context.Background())
	rctx.Next()
	return nil
}

func LogRequest(rctx *routing.Context) error {
	ctx := rctx.Get("ctx").(context.Context)
	ctx.With(zap.String("req", rctx.String())).Info("req")
	rctx.Next()
	return nil
}
