package postgres

import (
	"github.com/behouba/mediateq/database/schema"
	"github.com/behouba/mediateq/pkg/config"
)

// NewDatabase create a new postgres database
func NewDatabase(cfg *config.Database) (*schema.Database, error) {

	mediaTable, err := newMediaTable(nil)
	if err != nil {
		return nil, err
	}

	return &schema.Database{
		MediaTable: mediaTable,
	}, nil
}
