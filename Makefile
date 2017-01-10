default: install test build

install:
	glide install

test:
	go test ./...

build:
	go build -o bin/coffer ./
