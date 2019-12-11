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
	repository "url-at-minimal-api/internal/adapters/repository/redis"
	"url-at-minimal-api/internal/adapters/router"
	"url-at-minimal-api/internal/features/minifyurl"
	"url-at-minimal-api/internal/features/redirecturl"

	"github.com/go-redis/redis"
)

func main() {
	repository := repository.New(getRedisInstance())
	router := router.New(
		minify.New(minifyurl.New(repository, randomizer.New(clock.New()))),
		redirect.New(redirecturl.New(repository)),
		[]router.Middleware{middleware.Security},
	)

	println("I'm up!")
	log.Fatal(http.ListenAndServe(getPort(), router.Handler()))
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
