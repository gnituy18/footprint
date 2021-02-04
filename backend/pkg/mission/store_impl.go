package mission

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"footprint/pkg/context"
	"footprint/pkg/db"
)

func NewStore() Store {
	return &impl{
		col: db.MustGetMongo().Database("footprint").Collection("mission"),
	}
}

type impl struct {
	col *mongo.Collection
}

func (im *impl) Get(ctx context.Context, id string) (m *Mission, err error) {
	if err := im.col.FindOne(ctx, bson.M{"id": id}).Decode(&m); err != nil {
		ctx.With(zap.String("id", id)).Error("col.FindOne filed in mission.Store.Get")
		return nil, err
	}

	return
}

func (im *impl) Create(ctx context.Context, m *Mission) error {
	if _, err := im.col.InsertOne(ctx, m); err != nil {
		ctx.With(zap.Object("mission", m)).Error("col.InsertOne failed in mission.Store.Create")
		return err
	}

	return nil
}
