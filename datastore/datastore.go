package datastore

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/dty1er/hist-datastore/entity"
)

// Datastore is a GCP Cloud Datastore client
// which implements store interface
type Datastore struct {
	Cl *datastore.Client
}

func New() *Datastore {
	crd := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if crd == "" {
		log.Fatalf("gcp credential missing. set $GOOGLE_APPLICATION_CREDENTIALS")
	}
	ctx := context.Background()

	c, err := datastore.NewClient(ctx, "hist-datastore")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	cl := &Datastore{Cl: c}

	return cl
}

// GetAll gets records
func (cl *Datastore) GetAll(ctx context.Context) (hists entity.Histories, err error) {
	query := datastore.NewQuery("History").DistinctOn("Pwd", "Command").Order("-Timestamp").Limit(40000)
	_, err = cl.Cl.GetAll(ctx, query, &hists)
	if err != nil {
		return
	}
	return
}

// Put puts records
func (cl *Datastore) Put(ctx context.Context, pwd, cmd string) error {
	key := datastore.IncompleteKey("History", nil)
	hist := &entity.History{Pwd: pwd, Timestamp: time.Now(), Command: cmd}
	if _, err := cl.Cl.Put(ctx, key, hist); err != nil {
		return err
	}
	return nil
}
