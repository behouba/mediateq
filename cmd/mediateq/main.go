package main

import (
	"log"
	"net/http"

	"github.com/behouba/mediateq/routing"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	routing.Setup(router)

	log.Printf("Starting mediateq %s server at %s", "v0", ":8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
