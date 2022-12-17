package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/behouba/mediateq/pkg/config"
	"github.com/behouba/mediateq/routing"
	"github.com/behouba/mediateq/storage"
)

var (
	configFileFlag = flag.String("config", "mediateq.yaml", "The configuration file for mediateq server.")
)

func main() {

	// Parse the command line arguments
	flag.Parse()

	// Load configuration
	cfg, err := config.Load(*configFileFlag)
	if err != nil {
		log.Fatal(err)
	}

	// db, err := database.NewDatabase(cfg.Database)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	storage, err := storage.New(cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}

	handler, err := routing.NewHandler(cfg, storage, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), handler); err != nil {
		log.Fatal(err)
	}
}
