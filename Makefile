.PHONY: clean bin build build_cli build_srv vendor deploy

build: build_cli build_srv

clean:
	rm -rf bin/
	rm srv/dist/app.js

bin:
	mkdir -p bin

build_cli: bin
	go build -o bin/parrotify-hd-cli ./cli

build_srv: bin
	cd srv && make build

