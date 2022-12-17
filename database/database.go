package database

import (
	"fmt"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/database/postgres"
	"github.com/behouba/mediateq/database/schema"
	"github.com/behouba/mediateq/pkg/config"
)

func NewDatabase(cfg *config.Database) (*schema.Database, error) {
	if cfg.Type == mediateq.DBTypePostgres {
		return postgres.NewDatabase(cfg)
	}

	return nil, fmt.Errorf("database type %s is not supported", cfg.Type)
}
