package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/lgcavalheiro/go-svelte-embed-starter/frontend"
	"github.com/lgcavalheiro/go-svelte-embed-starter/services"
)

var apiPrefix = "/api"

func RegisterWebRoutes(e *echo.Echo) {
	e.GET("/*", echo.WrapHandler(frontend.DirFS))
}

func RegisterApiRoutes(e *echo.Echo) {
	e.GET(fmt.Sprintf("%s/healthcheck", apiPrefix), services.Healthcheck)
	e.GET(fmt.Sprintf("%s/double", apiPrefix), services.Double)
}
