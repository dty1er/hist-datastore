package dynamodb

import (
	"context"

	"github.com/dty1er/hist-datastore/entity"
)

// DynamoDB is a Amazon DynamoDB client
// which implements store interface
type DynamoDB struct {
}

func New() *DynamoDB {
	panic("implement me")
}

func (d *DynamoDB) Get(ctx context.Context, pwd string) ([]*entity.History, error) {
	panic("implement me")
}

// Put puts records
func (d *DynamoDB) Put(ctx context.Context, pwd, cmd string) error {
	panic("implement me")
}
