default: install test build

deps: update install

update:
	glide update
	
install:
	glide install

test:
	go test ./...

build:
	go build -o bin/coffer ./
