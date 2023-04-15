package main

import (
	"log"
	"net/http"

	"github.com/lgcavalheiro/go-svelte-embed-starter/routes"
)

func main() {
	routes.RegisterWebRoutes()
	routes.RegisterApiRoutes()

	log.Println("Listening @ http://localhost:1323")
	log.Fatal(http.ListenAndServe(":1323", nil))
}
