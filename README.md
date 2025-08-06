> [!WARNING]
> This is a work in progress!  It works, but it's a rough experience. Patches welcomed!

# HD Parrot Gif Maker

![A very basic UI for composing parrot gifs](.github/readme/screenshot.png)

This tool lets you compose [party parrot](https://cultofthepartyparrot.com) gifs using a web interface.

## Building

To build this project, you will need Go, Make, Node and Yarn installed.

```bash
$ make
mkdir -p bin
go build -o bin/parrotify-hd-cli ./cli
cd srv && make build
cd frontend && yarn build
yarn run v1.22.22
$ vite build
vite v2.9.18 building for production...
✓ 37 modules transformed.
dist/index.html                  0.44 KiB
dist/assets/index.59e1be2f.css   0.98 KiB / gzip: 0.45 KiB
dist/assets/index.c2bb5977.js    135.36 KiB / gzip: 45.26 KiB
✨  Done in 0.98s.
go build -tags=production -o ../bin/parrotify-hd-server .
$
```

## Running

To run the server, start `bin/parrotify-hd-server` and then open `http://localhost:3333` in your browser.
