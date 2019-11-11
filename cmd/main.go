package main

import (
	"log"
	"net/http"
	"url-at-minimal-api/internal/adapters/clock"
	"url-at-minimal-api/internal/adapters/handlers/minify"
	"url-at-minimal-api/internal/adapters/handlers/redirect"
	"url-at-minimal-api/internal/adapters/randomizer"
	"url-at-minimal-api/internal/adapters/repository"
	"url-at-minimal-api/internal/adapters/router"
	"url-at-minimal-api/internal/features/minifyurl"
	"url-at-minimal-api/internal/features/redirecturl"
)

func main() {
	repository := repository.New()
	router := router.New(
		minify.New(minifyurl.New(repository, randomizer.New(clock.New()))),
		redirect.New(redirecturl.New(repository)),
	)

	println("I'm up!")
	log.Fatal(http.ListenAndServe(":8080", router.Handler()))
}
