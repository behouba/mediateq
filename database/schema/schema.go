package schema

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/behouba/mediateq"
)

// Database represents mediateq database and group all database operations
type Database struct {
	MediaTable MediaTable
}

// MediaTable is an interface to represent the database operations on media objects
type MediaTable interface {
	Insert(ctx context.Context, media *mediateq.Media) (int64, error)
	SelectByHash(ctx context.Context, id string) (*mediateq.Media, error)
	// Get paginated list of medias
	SelectList(ctx context.Context, offset, limit int64) ([]mediateq.Media, error)
	Delete(ctx context.Context, id string) error
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
