package file

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dty1er/gtsv"
	"github.com/dty1er/hist-datastore/entity"
)

type File struct {
	path string
}

func New() *File {
	cache := fmt.Sprintf("%s/.cache/", os.Getenv("HOME"))

	xdgCacheHome := os.Getenv("XDG_CACHE_HOME")
	if xdgCacheHome != "" {
		cache = fmt.Sprintf("%s/", xdgCacheHome)
	}
	return &File{path: filepath.Join(cache, "hist")}
}

func (f *File) Get(ctx context.Context, pwd string) ([]*entity.History, error) {
	fi, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}

	defer fi.Close()

	hists := []*entity.History{}
	gt := gtsv.New(fi)
	for gt.Next() {
		cpwd := gt.String()
		ccommand := gt.String()
		_ = gt.String() // timestamp is not needed
		if cpwd == pwd {
			hists = append(hists, &entity.History{Command: ccommand, Pwd: pwd})
		}
	}

	return hists, nil
}

func (f *File) Put(ctx context.Context, pwd, cmd string) error {
	fi, err := os.OpenFile(f.path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fi.Close()
	fmt.Fprintln(fi, fmt.Sprintf("%s\t%s\t%s", pwd, cmd, time.Now().Format("2006-01-02 15:04:05")))

	return nil
}
