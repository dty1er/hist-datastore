package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/dty1er/hist-datastore/entity"
)

// DataStore ...
type DataStore interface {
	Get(ctx context.Context, pwd string)
	GetAll(ctx context.Context)
	Put(ctx context.Context, pwd, cmd string)
}

// Client ...
type Client struct {
	Cl *datastore.Client
}

// Get gets records
func (cl *Client) Get(ctx context.Context, pwd string) (hists entity.Histories, err error) {
	query := datastore.NewQuery("History").DistinctOn("Command").Filter("Pwd = ", pwd).Order("-Timestamp").Limit(5000)
	_, err = cl.Cl.GetAll(ctx, query, &hists)
	if err != nil {
		return
	}
	return
}

// GetAll gets records
func (cl *Client) GetAll(ctx context.Context) (hists entity.Histories, err error) {
	query := datastore.NewQuery("History").DistinctOn("Pwd", "Command").Order("-Timestamp").Limit(20000)
	_, err = cl.Cl.GetAll(ctx, query, &hists)
	if err != nil {
		return
	}
	return
}

// Put puts records
func (cl *Client) Put(ctx context.Context, pwd, cmd string) error {
	key := datastore.IncompleteKey("History", nil)
	hist := &entity.History{Pwd: pwd, Timestamp: time.Now(), Command: cmd}
	if _, err := cl.Cl.Put(ctx, key, hist); err != nil {
		return err
	}
	return nil
}
