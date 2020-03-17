package store

import (
	"context"

	"github.com/dty1er/hist-datastore/entity"
)

// Store is an interface to access to database
// storing history data
type Store interface {
	GetAll(ctx context.Context) (entity.Histories, error)
	Put(ctx context.Context, pwd, cmd string) error
}
