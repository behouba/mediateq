package postgres

import (
	"context"
	"database/sql"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/database/schema"
)

const (
	insertMediaSQL       = "INSERT INTO media (hash, content_type, origin, url, timestamp, size) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;"
	selectMediaByHashSQL = "SELECT id, hash, content_type, origin, url, timestamp, size FROM media WHERE hash=$1;"
	deleteMediaSQL       = ""
)

type mediaStmts struct {
	insertStmt     *sql.Stmt
	selectByIDStmt *sql.Stmt
	deleteStmt     *sql.Stmt
}

func newMediaTable(db *sql.DB) (schema.MediaTable, error) {
	s := &mediaStmts{}
	return s, schema.StatementList{
		{&s.insertStmt, insertMediaSQL},
		{&s.selectByIDStmt, selectMediaByHashSQL},
		// {&s.deleteStmt, deleteMediaSQL},
	}.Prepare(db)
}

// Delete implements schema.MediaTable
func (mediaStmts) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// SelectList implements schema.MediaTable
func (*mediaStmts) SelectList(ctx context.Context, offset int64, limit int64) ([]mediateq.Media, error) {
	panic("unimplemented")
}

// Insert implements schema.MediaTable
func (s mediaStmts) Insert(ctx context.Context, m *mediateq.Media) (int64, error) {
	var nid int64
	err := s.insertStmt.QueryRowContext(
		ctx, m.Hash, m.ContentType, m.Origin, m.URL, m.Timestamp, m.Size,
	).Scan(&nid)
	if err != nil {
		return 0, err
	}
	return nid, nil
}

// SelectByUID implements schema.MediaTable
func (m mediaStmts) SelectByHash(ctx context.Context, hash string) (*mediateq.Media, error) {
	md := mediateq.Media{}
	err := m.selectByIDStmt.QueryRowContext(ctx, hash).Scan(
		&md.ID, &md.Hash, &md.ContentType, &md.Origin, &md.URL, &md.Timestamp, &md.Size,
	)
	return &md, err
}
