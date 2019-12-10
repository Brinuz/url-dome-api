package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-at-minimal-api/internal/adapters/clock"
	"url-at-minimal-api/internal/adapters/handlers/minify"
	"url-at-minimal-api/internal/adapters/handlers/redirect"
	"url-at-minimal-api/internal/adapters/middleware"
	"url-at-minimal-api/internal/adapters/randomizer"
	repository "url-at-minimal-api/internal/adapters/repository/memory"
	"url-at-minimal-api/internal/adapters/router"
	"url-at-minimal-api/internal/features/minifyurl"
	"url-at-minimal-api/internal/features/redirecturl"
)

func main() {
	repository := repository.New()
	router := router.New(
		minify.New(minifyurl.New(repository, randomizer.New(clock.New()))),
		redirect.New(redirecturl.New(repository)),
		[]router.Middleware{middleware.Security},
	)

	println("I'm up!")
	log.Fatal(http.ListenAndServe(getPort(), router.Handler()))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}
