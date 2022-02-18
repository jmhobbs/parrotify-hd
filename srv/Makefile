.PHONY: dev build

dev:
	overmind start -f Procfile.dev

build:
	cd frontend && yarn build
	go build -tags=production -o ../bin/parrotify-hd-server .

