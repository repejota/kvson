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

dev-deps:


dist-clean: clean
	rm -rf pkg src bin
