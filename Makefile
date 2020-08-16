all:
	go build ./...

clean:
	rm hist-datastore

install:
	go install github.com/dty1er/hist-datastore/cmd/histdatastore
	go install github.com/dty1er/hist-datastore/cmd/cacheupdate
