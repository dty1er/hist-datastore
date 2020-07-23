package main

import (
	"log"

	"github.com/dty1er/hist-datastore/cache"
)

func main() {
	hists, err := cache.GetAll()
	if err != nil {
		log.Printf("Failed to get from datastore: %v", err)
	}

	err = cache.Update(hists)
	if err != nil {
		log.Printf("Failed to update cache: %v", err)
	}
}
