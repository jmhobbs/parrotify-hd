.PHONY: clean bin build build_cli build_srv build_srv_js

build: build_cli build_srv

clean:
	rm -rf bin/
	rm srv/dist/app.js

bin:
	mkdir -p bin

build_cli: bin
	go build -o bin/parrotify-hd-cli ./cli

build_srv: bin build_srv_js
	cd srv && make build

