package main

import (
	"net/http"

	"github.com/behouba/stash/image"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	image.Setup(router)

	http.ListenAndServe(":8080", router)
}
