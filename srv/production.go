// +build production

package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"os"
)

//go:embed frontend/dist
var distFS embed.FS

func logSink() io.Writer {
	return os.Stderr
}

func rootHandler() http.Handler {
	dist, err := fs.Sub(distFS, "frontend/dist")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(dist))
}
