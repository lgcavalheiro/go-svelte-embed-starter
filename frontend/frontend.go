package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var fsys embed.FS

func GetEmbeddedFiles() http.Handler {
	html, _ := fs.Sub(fsys, "dist")
	return http.FileServer(http.FS(html))
}
