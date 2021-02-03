package api

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"footprint/pkg/mission"
)

func Mount() {
	root := router.New()

	root.GET("/health", func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("healthy")
	})

	missionStore := mission.NewStore()
	missionGroup := root.Group("/mission")
	MountMissionRoutes(missionGroup, missionStore)

	fasthttp.ListenAndServe(":8080", root.Handler)
}
