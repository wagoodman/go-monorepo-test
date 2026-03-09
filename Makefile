.PHONY: build test clean release release-snapshot

build:
	go build -o dist/test ./cmd/test

test:
	go test -v ./...

clean:
	rm -rf dist/

release:
	goreleaser release --clean

release-snapshot:
	goreleaser release --snapshot --clean
