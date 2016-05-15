deps:
	go get github.com/constabulary/gb/...
	gb vendor restore
	npm i
	npm i -g electron-packager

build: clean deps
	gb test
	gb build

	gulp build

	rm -rf node_modules
	npm i --production

	cp bin/service dist/coffer/service
	cp package.json dist/
	cp -r node_modules dist/

pack: build
	electron-packager dist/ --platform=linux --arch=x64 --asar

clean:
	rm -rf coffer-*
