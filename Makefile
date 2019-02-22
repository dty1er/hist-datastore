all:
	go build ./...

clean:
	rm hist-datastore

install:
	go install github.com/yagi5/hist-datastore/cmd/histdatastore
	go install github.com/yagi5/hist-datastore/cmd/cacheupdate
