package main

import (
	"context"
	"log"

	"github.com/dty1er/hist-datastore/cache"
	"github.com/dty1er/hist-datastore/dynamodb"
)

func main() {
	store := dynamodb.New()

	hists, err := store.GetAll(context.Background())
	if err != nil {
		log.Printf("Failed to get from datastore: %v", err)
	}

	err = cache.Update(hists)
	if err != nil {
		log.Printf("Failed to update cache: %v", err)
	}
}
