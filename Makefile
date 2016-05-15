deps:
	go get github.com/constabulary/gb/...
	gb vendor restore
	npm i
	npm i -g electron-packager

build:
	gb test
	gb build

pack: clean build
	gulp build
	cp bin/service dist/coffer/service
	cp package.json dist/
	cp -r node_modules dist/

	electron-packager dist/ --platform=linux --arch=x64 --asar

clean:
	rm -rf dist/
	rm -rf coffer-*
