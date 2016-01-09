build:
	go build

install:
	go install

test:
	go test -v -race  ./... -coverprofile=coverage.out

cover:
	goveralls -v -coverprofile=coverage.out -service=circle-ci -repotoken=aAkZEnv0NuO7vPkra5A0ftypppCd3uIDQ
	rm coverage.out

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
