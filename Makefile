deps:
	go get github.com/constabulary/gb/...
	gb vendor restore
	npm i
	npm i -g electron-packager

test:
	gb test

build: clean deps
	gulp build

	rm -rf node_modules
	npm i --production

	cp -r node_modules dist/
	cp package.json dist/

build-linux:
	GOOS=linux GOARCH=amd64 gb build -P 1 -f -F
	cp bin/service-linux-amd64 dist/coffer/service
	electron-packager dist/ --platform=linux --arch=x64 --asar

build-win:
	GOOS=windows GOARCH=amd64 gb build -P 1 -f -F
	cp bin/service-windows-amd64.exe dist/coffer/service.exe
	electron-packager dist/ --platform=win32 --arch=x64 --asar

clean:
	rm -rf coffer-*
	rm -rf dist/
