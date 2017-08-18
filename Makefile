deps:
	rm -rf node_modules
	npm i
	npm i -g electron-packager electron

	dep ensure

test:
	go test github.com/mnzt/coffer

build:
    go build github.com/mnzt/coffer/cmd/coffer

clean:
	rm -rf coffer-*
	rm -rf dist/
