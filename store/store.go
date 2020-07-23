package store

import (
	"context"

	"github.com/dty1er/hist-datastore/entity"
)

// Store is an interface to access to database
// storing history data
type Store interface {
	Get(ctx context.Context, pwd string) ([]*entity.History, error)
	Put(ctx context.Context, pwd, cmd string) error
}
