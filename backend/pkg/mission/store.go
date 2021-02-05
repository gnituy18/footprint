package mission

import (
	"errors"

	"footprint/pkg/context"
)

var (
	ErrNotFound = errors.New("mission not found")
)

type Store interface {
	Get(ctx context.Context, ID string) (*Mission, error)
	Create(ctx context.Context, m *Mission) error
}
