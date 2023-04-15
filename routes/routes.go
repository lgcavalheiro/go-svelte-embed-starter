package routes

import (
	"fmt"
	"net/http"

	"github.com/lgcavalheiro/go-svelte-embed-starter/frontend"
	"github.com/lgcavalheiro/go-svelte-embed-starter/services"
)

var apiPrefix = "/api"

func RegisterWebRoutes() {
	http.Handle("/", frontend.GetEmbeddedFiles())
}

func RegisterApiRoutes() {
	http.HandleFunc(fmt.Sprintf("%s/healthcheck", apiPrefix), services.Healthcheck)
	http.HandleFunc(fmt.Sprintf("%s/double", apiPrefix), services.Double)
}
