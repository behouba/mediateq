package database

import (
	"fmt"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/database/postgres"
)

type dbType string

const (
	DBTypePostgres dbType = "postgres"
	DBTypeSQLite   dbType = "sqlite"
)

func NewDatabase(cfg *mediateq.DBConfig, dbType string) (mediateq.Database, error) {
	if dbType == "postgresql" {
		return postgres.NewDatabase(cfg)
	}

	return nil, fmt.Errorf("database type %s is not supported", dbType)
}
