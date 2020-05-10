package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"url-at-minimal-api/internal/external_interfaces/clock"
	"url-at-minimal-api/internal/external_interfaces/handlers/minify"
	"url-at-minimal-api/internal/external_interfaces/handlers/redirect"
	"url-at-minimal-api/internal/external_interfaces/middleware"
	"url-at-minimal-api/internal/external_interfaces/randomizer"
	repository "url-at-minimal-api/internal/external_interfaces/repository/postgres"
	"url-at-minimal-api/internal/external_interfaces/rest"
	"url-at-minimal-api/internal/use_cases/minifyurl"
	"url-at-minimal-api/internal/use_cases/redirecturl"
)

func main() {
	repository := repository.New(getPostgresInstance())
	rest := rest.New(
		minify.New(minifyurl.New(repository, randomizer.New(clock.New()))),
		redirect.New(redirecturl.New(repository)),
		[]rest.Middleware{middleware.Security},
	)

	println("I'm up!")
	log.Fatal(http.ListenAndServe(getPort(), rest.Handler()))
}

func getPostgresInstance() *sql.DB {
	host := getEnvDefault("POSTGRES_HOST", "localhost")
	port := getEnvDefault("POSTGRES_PORT", "2345")
	user := getEnvDefault("POSTGRES_USER", "postgres")
	password := getEnvDefault("POSTGRES_PASSWORD", "secretpassword")
	dbname := getEnvDefault("POSTGRES_DB", "url-dome")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to get postgres instance.")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to postgres instance.")
	}

	return db
}

func getEnvDefault(key, defaulValue string) string {
	value, hasValue := os.LookupEnv(key)
	if !hasValue {
		return defaulValue
	}
	return value
}

func getPort() string {
	return fmt.Sprintf(":%s", getEnvDefault("PORT", "8080"))
}
