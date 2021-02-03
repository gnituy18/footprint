package api

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"footprint/pkg/mission"
)

func MountMissionRoutes(group *router.Group, missionStore mission.Store) {
	handler := &missionHandler{
		missionStore: missionStore,
	}

	group.GET("/{ID}", handler.getMission)
}

type missionHandler struct {
	missionStore mission.Store
}

func (mh *missionHandler) getMission(ctx *fasthttp.RequestCtx) {
	fmt.Println("hello world")
}
