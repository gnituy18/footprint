package api

import (
	"github.com/qiangxue/fasthttp-routing"

	"footprint/pkg/log"
	"footprint/pkg/mission"
)

func Router() *routing.Router {
	root := routing.New()
	root.Use(WithContext)

	root.Get("/health", func(rctx *routing.Context) error {
		rctx.WriteString("healthy")
		log.Global().Info("healthy")
		return nil
	})

	missionStore := mission.NewStore()
	missionGroup := root.Group("/mission")
	MountMissionRoutes(missionGroup, missionStore)

	return root
}
