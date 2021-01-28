package mission

import (
	"footprint/pkg/context"
)

type Store interface {
	Get(ctx context.Context, ID string) (*Mission, error)
	Create(ctx context.Context, m *Mission) error
}
