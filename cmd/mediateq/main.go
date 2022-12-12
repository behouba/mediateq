package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/behouba/mediateq/config"
	"github.com/behouba/mediateq/routing"
)

var (
	configFileFlag = flag.String("config", "mediateq.yaml", "This flag is used to specify the location of the configuration file for the application. The default value for this flag is \"mediateq.yaml\".")
)

func main() {

	// Parse the command line arguments
	flag.Parse()

	// Load configuration
	cfg, err := config.Load(*configFileFlag)
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
