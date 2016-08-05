deps:
	rm -rf node_modules
	npm i --production

	go get github.com/constabulary/gb/...
	gb vendor restore
	npm i
	npm i -g electron-packager

test:
	gb test

build-app: clean
	gulp build
	gb build

	cp package.json dist/coffer/

build-linux: build-app
	GOOS=linux GOARCH=amd64 gb build -P 1 -f -F
	cp bin/service-linux-amd64 dist/coffer/service
	electron-packager dist/coffer --platform=linux --arch=x64 --asar

build-win: build-app
	GOOS=windows GOARCH=amd64 gb build -P 1 -f -F
	cp bin/service-windows-amd64.exe dist/coffer/service.exe
	electron-packager dist/coffer --platform=win32 --arch=x64 --asar

clean:
	rm -rf coffer-*
	rm -rf dist/
