package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
)

const (
	projectID = "hist-datastore"
)

// DataStore ...
type DataStore interface {
	Get(ctx context.Context, pwd string)
	Put(ctx context.Context, pwd, cmd string)
}

// History represents the kind of Datastore
type History struct {
	Command   string
	Pwd       string
	Timestamp time.Time
}

// Client ...
type Client struct {
	ds *datastore.Client
}

func main() {
	ctx := context.Background()

	c, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	cl := &Client{ds: c}

	op, pwd, cmd := os.Args[1], os.Args[2], os.Args[3:]

	if op != "put" && op != "get" {
		log.Fatalf("specify get or put")
	}

	// Put
	if op == "put" {
		go func() {
			if err := cl.Put(ctx, pwd, strings.Join(cmd, " ")); err != nil {
				log.Fatalf("Failed to save history: %v", err)
			}
		}()
		return
	}

	// Get
	var hists []History
	hists, err = getFromCache(pwd)
	if err != nil {
		log.Printf("Failed to get from cache: %v", err)

		hists, err = cl.Get(ctx, pwd)
		if err != nil {
			log.Printf("Failed to get history: %v", err)
		}
	}
	printHists(hists)
}

func printHists(hists []History) {
	for _, hist := range hists {
		fmt.Println(hist.Command)
	}
}

func getFromCache(pwd string) (hists []History, err error) {
	cacheFileName := "hist-datastore"
	cache := fmt.Sprintf("%s/.cache/%s", os.Getenv("HOME"), cacheFileName)

	xdgCacheHome := os.Getenv("XDG_CACHE_HOME")
	if xdgCacheHome != "" {
		cache = fmt.Sprintf("%s/%s", xdgCacheHome, cacheFileName)
	}
	cacheFile, err := os.Open(cache)
	if err != nil {
		return
	}
	cacheBytes, err := ioutil.ReadAll(cacheFile)
	if err != nil {
		return
	}

	// cache file format
	//

	return nil, nil
}

// Get ...
func (cl *Client) Get(ctx context.Context, pwd string) (hists []History, err error) {
	query := datastore.NewQuery("History").Filter("Pwd = ", pwd).Limit(100)
	_, err = cl.ds.GetAll(ctx, query, &hists)
	if err != nil {
		return
	}
	return
}

// Put ...
func (cl *Client) Put(ctx context.Context, pwd, cmd string) error {
	key := datastore.IncompleteKey("History", nil)
	hist := &History{Timestamp: time.Now(), Pwd: pwd, Command: cmd}
	if _, err := cl.ds.Put(ctx, key, hist); err != nil {
		return err
	}
	return nil
}
