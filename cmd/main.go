package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-at-minimal-api/internal/external_interfaces/clock"
	"url-at-minimal-api/internal/external_interfaces/handlers/minify"
	"url-at-minimal-api/internal/external_interfaces/handlers/redirect"
	"url-at-minimal-api/internal/external_interfaces/middleware"
	"url-at-minimal-api/internal/external_interfaces/randomizer"
	repository "url-at-minimal-api/internal/external_interfaces/repository/redis"
	"url-at-minimal-api/internal/external_interfaces/rest"
	"url-at-minimal-api/internal/use_cases/minifyurl"
	"url-at-minimal-api/internal/use_cases/redirecturl"

	"github.com/go-redis/redis"
)

func main() {
	repository := repository.New(getRedisInstance())
	rest := rest.New(
		minify.New(minifyurl.New(repository, randomizer.New(clock.New()))),
		redirect.New(redirecturl.New(repository)),
		[]rest.Middleware{middleware.Security},
	)

	println("I'm up!")
	log.Fatal(http.ListenAndServe(getPort(), rest.Handler()))
}

func getRedisInstance() *redis.Client {
	rURL := os.Getenv("REDIS_URL")
	if rURL == "" {
		rURL = "redis://localhost:6379"
	}

	opts, err := redis.ParseURL(rURL)
	if err != nil {
		log.Fatal("Failed to get redis instance.")
	}

	client := redis.NewClient(opts)
	pong, err := client.Ping().Result()
	if err != nil || pong != "PONG" {
		log.Fatal("Failed to connect to redis instance.")
	}
	return client
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}
