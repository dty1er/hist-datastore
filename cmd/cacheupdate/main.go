package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/dty1er/hist-datastore/cache"
	histds "github.com/dty1er/hist-datastore/datastore"
)

const (
	projectID = "hist-datastore"
)

func main() {
	crd := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if crd == "" {
		log.Fatalf("gcp credential missing. set $GOOGLE_APPLICATION_CREDENTIALS")
	}
	ctx := context.Background()

	c, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	cl := &histds.Client{Cl: c}

	// Get
	hists, err := cl.GetAll(ctx)
	if err != nil {
		log.Printf("Failed to get from datastore: %v", err)
	}
	err = cache.Update(hists)
	if err != nil {
		log.Printf("Failed to update cache: %v", err)
	}
}
