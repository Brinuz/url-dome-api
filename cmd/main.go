package main

import (
	"log"
	"net/http"
	"url-at-minimal-api/internal/adapters/clock"
	"url-at-minimal-api/internal/adapters/handlers/minify"
	"url-at-minimal-api/internal/adapters/randomizer"
	"url-at-minimal-api/internal/adapters/repository"
	"url-at-minimal-api/internal/adapters/router"
	"url-at-minimal-api/internal/features/minifyurl"
)

func main() {
	router := router.New(minify.New(minifyurl.New(repository.New(), randomizer.New(clock.New()))))

	println("I'm up!")
	log.Fatal(http.ListenAndServe(":8080", router.Handler()))
}
