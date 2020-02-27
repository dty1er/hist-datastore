package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dty1er/gtsv"
	"github.com/dty1er/hist-datastore/entity"
)

// Cache ...
type Cache interface {
	Get(pwd string)
	GetAll()
}

func getCacheDir() string {
	cache := fmt.Sprintf("%s/.cache/", os.Getenv("HOME"))

	xdgCacheHome := os.Getenv("XDG_CACHE_HOME")
	if xdgCacheHome != "" {
		cache = fmt.Sprintf("%s/", xdgCacheHome)
	}
	return cache
}

func getCacheFileName() string {
	cacheFileName := "hist-datastore"
	return getCacheDir() + cacheFileName
}

// Update ...
func Update(hists entity.Histories) error {
	cacheDir := getCacheDir()
	oldCache := getCacheFileName()
	newCache := getCacheFileName()
	// Backup
	err := os.Rename(oldCache, fmt.Sprintf("%s%s%s", cacheDir, "hist-datastore", time.Now().Format(time.RFC3339)))
	if err != nil {
		return err
	}
	newCacheFile, err := os.OpenFile(newCache, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer newCacheFile.Close()
	for _, hist := range hists {
		fmt.Fprintln(newCacheFile, fmt.Sprintf("%s\t%s\t%s", hist.Pwd, hist.Command, time.Now().Format("2006-01-02 15:04:05")))
	}

	return nil
}

// Put puts cache
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

// Get gets histories by pwd
func Get(pwd string) (hists entity.Histories, err error) {
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
			hists = append(hists, &entity.History{Command: ccommand, Pwd: pwd})
		}
	}
	return
}

// GetAll gets all histories
func GetAll() (hists entity.Histories, err error) {
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
		hists = append(hists, &entity.History{Command: ccommand, Pwd: cpwd})
	}
	return
}
