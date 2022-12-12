package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/behouba/mediateq/config"
	"github.com/behouba/mediateq/routing"
)

const ()

func main() {

	// Load configuration
	cfg, err := config.Load("mediateq.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// database, err := database.NewDatabase(cfg.Database, database.TypePostgres)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// storage, err := localdisk.Newstorage(cfg.Storage)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	handler, err := routing.NewHandler(cfg, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), handler); err != nil {
		log.Fatal(err)
	}
}
