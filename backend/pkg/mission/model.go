package mission

import (
	"go.uber.org/zap/zapcore"
)

type Mission struct {
	ID          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
}

func (m *Mission) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("id", m.ID)
	encoder.AddString("name", m.Name)
	encoder.AddString("description", m.Description)
	return nil
}
