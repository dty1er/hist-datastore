package datastore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
)

// History represents the kind of Datastore
type History struct {
	Pwd       string
	Command   string
	Timestamp time.Time
}

// Histories ...
type Histories []*History

// DataStore ...
type DataStore interface {
	Get(ctx context.Context, pwd string)
	Put(ctx context.Context, pwd, cmd string)
}

// Client ...
type Client struct {
	Cl *datastore.Client
}

// Print ...
func (hs Histories) Print() {
	for _, h := range hs {
		fmt.Println(h.Command)
	}
}

// Get ...
func (cl *Client) Get(ctx context.Context, pwd string) (hists Histories, err error) {
	query := datastore.NewQuery("History").Filter("Pwd = ", pwd).Order("-Timestamp").Limit(5000)
	_, err = cl.Cl.GetAll(ctx, query, &hists)
	if err != nil {
		return
	}
	return
}

// Put ...
func (cl *Client) Put(ctx context.Context, pwd, cmd string) error {
	key := datastore.IncompleteKey("History", nil)
	hist := &History{Pwd: pwd, Timestamp: time.Now(), Command: cmd}
	if _, err := cl.Cl.Put(ctx, key, hist); err != nil {
		return err
	}
	return nil
}
