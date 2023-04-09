package frontend

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed dist
var dir embed.FS

var DirFS = http.FileServer(http.FS(echo.MustSubFS(dir, "dist")))
