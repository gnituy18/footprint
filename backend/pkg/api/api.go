package api

import (
	"github.com/qiangxue/fasthttp-routing"

	"footprint/pkg/mission"
)

func Router() *routing.Router {
	root := routing.New()
	root.Use(WithContext)

	root.Get("/health", func(ctx *routing.Context) error {
		ctx.WriteString("healthy")
		return nil
	})

	missionStore := mission.NewStore()
	missionGroup := root.Group("/mission")
	MountMissionRoutes(missionGroup, missionStore)

	return root
}
