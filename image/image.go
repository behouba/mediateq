package image

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Setup(r *chi.Mux) {

	r.Post("/image", upload)

	r.Get("/hello", upload)
}

func upload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello image router`))
}
