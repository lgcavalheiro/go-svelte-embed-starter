package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/lgcavalheiro/go-svelte-embed-starter/routes"
)

func loadConfig(path string) {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		env := strings.SplitN(fileScanner.Text(), "=", 2)
		if env[0] == "" {
			continue
		}
		os.Setenv(env[0], env[1])
	}

	f.Close()
}

func main() {
	loadConfig(".env")

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}

	routes.RegisterWebRoutes()
	routes.RegisterApiRoutes()

	log.Printf("Listening @ http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
