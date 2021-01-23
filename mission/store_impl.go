package mission

import (
	"context"
)

func NewStore() Store {
	return &impl{}
}

type impl struct{}

func (im *impl) Get(ctx context.Context, ID string) (m *Mission, err error) {
	m = &Mission{
		ID:          ID,
		Name:        "Footprint",
		Description: "Accelerate mission completion",
	}

	return
}

func (im *impl) Create(ctx context.Context, m *Mission) (err error) {
	return
}
