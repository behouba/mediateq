package database

import (
	"fmt"

	"github.com/behouba/mediateq/config"
	"github.com/behouba/mediateq/database/postgres"
	"github.com/behouba/mediateq/database/schema"
)

type dbType string

const (
	TypePostgres dbType = "postgres"
	TypeSQLite   dbType = "sqlite"
)

func NewDatabase(cfg *config.Database, dbType dbType) (*schema.Database, error) {
	if dbType == TypePostgres {
		return postgres.NewDatabase(cfg)
	}

	return nil, fmt.Errorf("database type %s is not supported", dbType)
}
