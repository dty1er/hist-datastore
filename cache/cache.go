package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/yagi5/gtsv"
	"github.com/yagi5/hist-datastore/datastore"
)

func getCacheFileName() string {
	cacheFileName := "hist-datastore"
	cache := fmt.Sprintf("%s/.cache/%s", os.Getenv("HOME"), cacheFileName)

	xdgCacheHome := os.Getenv("XDG_CACHE_HOME")
	if xdgCacheHome != "" {
		cache = fmt.Sprintf("%s/%s", xdgCacheHome, cacheFileName)
	}
	return cache
}

// Put ...
func Put(ctx context.Context, pwd, cmd string) error {
	cache := getCacheFileName()
	file, err := os.OpenFile(cache, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, fmt.Sprintf("%s\t%s\t%s", pwd, cmd, time.Now().Format("2006-01-02 15:04:05")))

	return nil
}

// Get ...
func Get(pwd string) (hists datastore.Histories, err error) {
	cache := getCacheFileName()
	f, err := os.Open(cache)
	if err != nil {
		return
	}

	gt := gtsv.New(f)
	for gt.Next() {
		cpwd := gt.String()
		ccommand := gt.String()
		_ = gt.String() // timestamp is not needed
		if cpwd == pwd {
			hists = append(hists, &datastore.History{Command: ccommand, Pwd: pwd})
		}
	}
	return
}
