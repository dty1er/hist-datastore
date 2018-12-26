package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
)

const (
	projectID = "hist-datastore"
)

var client *datastore.Client

// History represents the kind of Datastore
type History struct {
	Command   string
	Pwd       string
	Timestamp time.Time
}

func init() {
	c, err := datastore.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	client = c
}

func main() {
	ctx := context.Background()
	op, pwd, cmd := os.Args[1], os.Args[2], os.Args[3:]
	if op != "put" && op != "get" {
		log.Fatalf("specify get or put")
	}

	if op == "put" {
		put(ctx, pwd, strings.Join(cmd, " "))
		return
	}
	get(ctx, pwd)
}

func get(ctx context.Context, pwd string) {
	var hists []History
	query := datastore.NewQuery("History").Filter("Pwd = ", pwd).Limit(100)
	_, err := client.GetAll(ctx, query, &hists)
	if err != nil {
		log.Fatalf("Failed to get history: %v", err)
	}
	for _, hist := range hists {
		fmt.Println(hist.Command)
	}
}

func put(ctx context.Context, pwd, cmd string) {
	if _, err := client.Put(ctx, datastore.IncompleteKey("History", nil),
		&History{Timestamp: time.Now(), Pwd: pwd, Command: cmd}); err != nil {
		log.Fatalf("Failed to save history: %v", err)
	}
}
