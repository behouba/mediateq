package schema

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/behouba/mediateq"
)

// Database represents mediateq database and group all database operations
type Database struct {
	MediaTable     MediaTable
	ThumbnailTable ThumbnailTable
}

// MediaTable is an interface to represent the database operations on media objects
type MediaTable interface {
	Insert(ctx context.Context, media *mediateq.Media) error
	SelectByBase64Hash(ctx context.Context, base64Hash string) (*mediateq.Media, error)
	SelectList(ctx context.Context, offset, limit int) ([]mediateq.Media, error)
	Delete(ctx context.Context, id string) error
}

type ThumbnailTable interface {
	Insert(ctx context.Context, thumbnail *mediateq.Thumbnail) error
	Select(ctx context.Context, mediaID string, width, height int, crop bool) (*mediateq.Thumbnail, error)
	SelectByMediaID(ctx context.Context, mediaID string) ([]mediateq.Thumbnail, error)
	Delete(ctx context.Context, mediaID string) error
}

// statementList is a list of SQL statements to prepare and a pointer to where to store the resulting prepared statement.
type StatementList []struct {
	Statement **sql.Stmt
	SQL       string
}

// Prepare the SQL for each statement in the list and assign the result to the prepared statement.
func (s StatementList) Prepare(db *sql.DB) (err error) {
	for _, statement := range s {
		if *statement.Statement, err = db.Prepare(statement.SQL); err != nil {
			err = fmt.Errorf("error %q while preparing statement: %s", err, statement.SQL)
			return
		}
	}
	return
}
