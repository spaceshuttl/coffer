deps:
	go get github.com/constabulary/gb/...
	npm i
	npm i -g electron-packager

build: deps
	gb test
	gb build

pack: clean build
	gulp build

	rm -rf node_modules
	npm i --production

	cp bin/service dist/coffer/service
	cp package.json dist/
	cp -r node_modules dist/

	electron-packager dist/ --platform=linux --arch=x64 --asar

clean:
	rm -rf dist/
	rm -rf coffer-*
