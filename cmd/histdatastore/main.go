package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dty1er/hist-datastore/entity"
	"github.com/dty1er/hist-datastore/file"
	"github.com/dty1er/hist-datastore/store"
)

func main() {
	// TODO: use flag.Parse and enable to specify database
	store := file.New()

	switch os.Args[1] {
	case "put":
		Put(store, os.Args[2], os.Args[3:])
	case "get":
		Get(store, os.Args[2])
	}
}

func Put(store store.Store, dir string, cmd []string) {
	ctx := context.Background()
	if err := store.Put(ctx, dir, strings.Join(cmd, " ")); err != nil {
		log.Fatalf("Failed to save history: %v", err)
	}
}

func Get(store store.Store, dir string) {
	hists, err := store.Get(context.Background(), dir)
	if err != nil {
		log.Fatalf("Failed to get histories: %v", err)
	}

	var uhs []*entity.History
	m := make(map[string]bool)
	for _, h := range hists {
		if !m[h.Command] {
			m[h.Command] = true
			uhs = append(uhs, h)
		}
	}
	for _, h := range uhs {
		fmt.Println(h.Command)
	}
}
