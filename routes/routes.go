package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lgcavalheiro/go-svelte-embed-starter/frontend"
	"github.com/lgcavalheiro/go-svelte-embed-starter/services"
)

func RegisterWebRoutes() {
	http.Handle("/", frontend.GetEmbeddedFiles())
}

func RegisterApiRoutes() {
	apiPrefix := os.Getenv("API_PREFIX")

	http.HandleFunc(fmt.Sprintf("%s/healthcheck", apiPrefix), services.Healthcheck)
	http.HandleFunc(fmt.Sprintf("%s/double", apiPrefix), services.Double)
}
