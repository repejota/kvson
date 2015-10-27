build:
	go build

install:
	go install

test:
	go test -v ./...

cover:
	go test -v ./... -cover

lint:
	go vet ./...
	golint ./...

clean:
	go clean

deps:
	# Dev dependencies

dist-clean: clean
	rm -rf pkg src bin

.PHONY: build test lint deps clean dist-clean
