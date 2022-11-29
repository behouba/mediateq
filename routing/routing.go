package routing

import (
	"net/http"
	"time"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const version = "v0"

type mux struct {
	cfg            *config.Config
	handler        http.Handler
	db             mediateq.Database
	storage        mediateq.FileStorage
	startTimestamp int64
}

func NewMux() *mux {

	r := chi.NewMux()

	Setup(r)

	return &mux{
		cfg:            &config.Config{},
		handler:        r,
		db:             nil,
		storage:        nil,
		startTimestamp: time.Now().Unix(),
	}
}

func Setup(r *chi.Mux) {

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	mux := mux{
		cfg:            &config.Config{},
		handler:        r,
		db:             nil,
		storage:        nil,
		startTimestamp: time.Now().Unix(),
	}

	r.Route("/mediateq/"+version, func(r chi.Router) {

		r.Get("/info", mux.infoHandler)
	})

}
