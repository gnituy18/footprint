package api

import (
	"fmt"

	"github.com/qiangxue/fasthttp-routing"

	"footprint/pkg/mission"
)

func MountMissionRoutes(group *routing.RouteGroup, missionStore mission.Store) {
	handler := &missionHandler{
		missionStore: missionStore,
	}

	group.Get("/<id>", handler.getMission)
}

type missionHandler struct {
	missionStore mission.Store
}

func (mh *missionHandler) getMission(ctx *routing.Context) error {
	fmt.Println("hello world")
	return nil
}
