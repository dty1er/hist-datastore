package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/dty1er/hist-datastore/cache"
	"github.com/dty1er/hist-datastore/dynamodb"
	"github.com/dty1er/hist-datastore/store"
)

func main() {
	store := dynamodb.New()
	// store := datastore.New()

	switch os.Args[1] {
	case "put":
		Put(store, os.Args[2], os.Args[3:])
	case "get":
		Get(os.Args[2])
	}
}

func Put(store store.Store, dir string, cmd []string) {
	ctx := context.Background()
	if err := store.Put(ctx, dir, strings.Join(cmd, " ")); err != nil {
		log.Fatalf("Failed to save history: %v", err)
	}

	cache.Put(ctx, dir, strings.Join(cmd, " "))
}

func Get(dir string) {
	hists, err := cache.Get(dir)
	if err != nil {
		log.Printf("Failed to get from cache: %v", err)
	}
	if len(hists) <= 30 {
		allHists, err := cache.GetAll()
		if err != nil {
			log.Printf("Failed to get from cache: %v", err)
		}
		hists = append(hists, allHists...)
	}
	hists.Print()
}
