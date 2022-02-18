// +build production

package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed frontend/dist
var distFS embed.FS

func rootHandler() http.Handler {
	dist, err := fs.Sub(distFS, "frontend/dist")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(dist))
}