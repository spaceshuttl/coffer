default: install test build

deps: update install

update:
	glide update

install:
	glide install

test:
	go test ./...

sign:
	drone sign spaceshuttl/coffer
	
build:
	go build -o bin/coffer ./
