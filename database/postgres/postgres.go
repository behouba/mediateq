package postgres

import (
	"database/sql"
	"fmt"

	"github.com/behouba/mediateq/database/schema"
	"github.com/behouba/mediateq/pkg/config"

	_ "github.com/lib/pq"
)

// NewDatabase create a new postgres database
func NewDatabase(cfg *config.Database) (*schema.Database, error) {

	db, err := sql.Open("postgres",
		fmt.Sprintf("dbname=%v user=%v password=%v port=%v", cfg.DBName, cfg.Username, cfg.Password, cfg.Port),
	)
	if err != nil {
		return nil, err
	}

	mediaTable, err := newMediaTable(db)
	if err != nil {
		return nil, err
	}

	thumbnailTable, err := newThumbnailTable(db)
	if err != nil {
		return nil, err
	}

	return &schema.Database{
		MediaTable:     mediaTable,
		ThumbnailTable: thumbnailTable,
	}, nil
}
