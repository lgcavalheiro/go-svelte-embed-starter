package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lgcavalheiro/go-svelte-embed-starter/routes"
)

func createServer() *echo.Echo {
	e := echo.New()
	routes.RegisterWebRoutes(e)
	routes.RegisterApiRoutes(e)
	return e
}

func main() {
	e := createServer()
	e.Logger.Fatal(e.Start(":1323"))
}
