package api

import (
	"encoding/json"
	"net/http"

	"github.com/qiangxue/fasthttp-routing"
	"go.uber.org/zap"

	"footprint/pkg/context"
	"footprint/pkg/mission"
)

func MountMissionRoutes(group *routing.RouteGroup, missionStore mission.Store) {
	handler := &missionHandler{
		missionStore: missionStore,
	}

	group.Post("/", handler.createMission)
	group.Get("/<id>", handler.getMission)
}

type missionHandler struct {
	missionStore mission.Store
}

func (mh *missionHandler) createMission(rctx *routing.Context) error {
	ctx := rctx.Get("ctx").(context.Context)

	body := rctx.Request.Body()

	m := &mission.Mission{}
	if err := json.Unmarshal(body, m); err != nil {
		ctx.With(
			zap.Error(err),
			zap.String("body", string(body)),
		).Error("missionHandler.missionStore.Get failed in missionHandler.getMission")
		return err
	}

	err := mh.missionStore.Create(ctx, m)
	if err != nil {
		ctx.With(zap.Error(err)).Error("missionHandler.missionStore.Create failed in missionHandler.getMission")
		return err
	}

	JSON(rctx, http.StatusOK, m)
	return nil
}

func (mh *missionHandler) getMission(rctx *routing.Context) error {
	ctx := rctx.Get("ctx").(context.Context)
	id := rctx.Param("id")

	m, err := mh.missionStore.Get(ctx, id)
	if err != nil {
		ctx.With(
			zap.Error(err),
			zap.String("id", id),
		).Error("missionHandler.missionStore.Get failed in missionHandler.getMission")
		return err
	}

	JSON(rctx, http.StatusOK, m)
	return nil
}
