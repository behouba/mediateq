package database

import (
	"fmt"

	"github.com/behouba/mediateq/database/postgres"
	"github.com/behouba/mediateq/database/schema"
	"github.com/behouba/mediateq/pkg/config"
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
