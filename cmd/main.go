package main

import (
	"log"
	"net/http"
	"url-at-minimal-api/internal/router"
)

func main() {
	router := router.Router{}

	println("I'm up!")
	log.Fatal(http.ListenAndServe(":8080", router.Handler()))
}
