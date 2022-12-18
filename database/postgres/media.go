package postgres

import (
	"context"
	"database/sql"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/database/schema"
)

const (
	insertMediaSQL = "INSERT INTO media () VALUES ()"
	selectMediaSQL = ""
	deleteMediaSQL = ""
)

type mediaStmts struct {
	insertSQL  *sql.Stmt
	selectStmt *sql.Stmt
	deleteStmt *sql.Stmt
}

func newMediaTable(db *sql.DB) (schema.MediaTable, error) {
	s := mediaStmts{}
	return s, schema.StatementList{
		{&s.insertSQL, insertMediaSQL},
		{&s.selectStmt, selectMediaSQL},
		{&s.deleteStmt, deleteMediaSQL},
	}.Prepare(db)
}

// Delete implements schema.MediaTable
func (mediaStmts) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Insert implements schema.MediaTable
func (mediaStmts) Insert(ctx context.Context, media *mediateq.Media) (int, error) {
	panic("unimplemented")
}

// SelectByUID implements schema.MediaTable
func (mediaStmts) SelectByID(ctx context.Context, id string) (*mediateq.Media, error) {
	panic("unimplemented")
}
