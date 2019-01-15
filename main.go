package main

import (
	"context"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/yagi5/hist-datastore/cache"
	histds "github.com/yagi5/hist-datastore/datastore"
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

	op := os.Args[1]

	if op != "put" && op != "get" {
		log.Fatalf("specify get or put")
	}

	// Put
	if op == "put" {
		pwd, cmd := os.Args[2], os.Args[3:]
		if err := cl.Put(ctx, pwd, strings.Join(cmd, " ")); err != nil {
			log.Fatalf("Failed to save history: %v", err)
		}
		cache.Put(ctx, pwd, strings.Join(cmd, " "))
		return
	}

	// Get
	pwd := os.Args[2]
	hists, err := cl.Get(ctx, pwd)
	if err != nil {
		log.Printf("Failed to get history: %v", err)
		hists, err = cache.Get(pwd)
		if err != nil {
			log.Printf("Failed to get from cache: %v", err)
		}
	}
	if len(hists) <= 10 {
		allHists, err := cache.GetAll()
		if err != nil {
			log.Printf("Failed to get from cache: %v", err)
		}
		hists = append(hists, allHists...)
	}
	hists.Print()
}
