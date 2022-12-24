package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/database/schema"
)

const (
	insertMediaSQL = `
	INSERT INTO 
		media (base64_hash, file_path, content_type, origin, url, timestamp, size) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7) 
	ON CONFLICT DO NOTHING
	RETURNING id;
	`

	selectMediaFormat = "SELECT id, base64_hash, file_path, content_type, origin, url, timestamp, size FROM media %s;"

	deleteMediaSQL = ""
)

var (
	selectMediaListSQL = fmt.Sprintf(selectMediaFormat, "OFFSET $1 LIMIT CASE WHEN $2 = 0 THEN 10 ELSE $2 END;")

	selectMediaByHashSQL = fmt.Sprintf(selectMediaFormat, "WHERE base64_hash=$1")
)

type mediaStmts struct {
	insertStmt       *sql.Stmt
	selectByHashStmt *sql.Stmt
	selectListStmt   *sql.Stmt
	deleteStmt       *sql.Stmt
}

func newMediaTable(db *sql.DB) (schema.MediaTable, error) {
	s := &mediaStmts{}
	return s, schema.StatementList{
		{&s.insertStmt, insertMediaSQL},
		{&s.selectByHashStmt, selectMediaByHashSQL},
		{&s.selectListStmt, selectMediaListSQL},
		{&s.deleteStmt, deleteMediaSQL},
	}.Prepare(db)
}

// Delete implements schema.MediaTable
func (mediaStmts) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// SelectList implements schema.MediaTable
func (s mediaStmts) SelectList(ctx context.Context, offset int, limit int) ([]mediateq.Media, error) {
	mediaList := make([]mediateq.Media, 0)

	rows, err := s.selectListStmt.QueryContext(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		m := mediateq.Media{}

		if err := rows.Scan(
			&m.ID, &m.Base64Hash, &m.FilePath, &m.ContentType,
			&m.Origin, &m.URL, &m.Timestamp, &m.Size,
		); err != nil {
			return nil, err
		}
		mediaList = append(mediaList, m)
	}

	return mediaList, nil
}

// Insert implements schema.MediaTable
func (s mediaStmts) Insert(ctx context.Context, m *mediateq.Media) (int64, error) {
	var id int64
	err := s.insertStmt.QueryRowContext(
		ctx, m.Base64Hash, m.FilePath, m.ContentType, m.Origin, m.URL, m.Timestamp, m.Size,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// SelectByUID implements schema.MediaTable
func (s mediaStmts) SelectByHash(ctx context.Context, hash string) (*mediateq.Media, error) {
	md := mediateq.Media{}
	err := s.selectByHashStmt.QueryRowContext(ctx, hash).Scan(
		&md.ID, &md.Base64Hash, &md.FilePath, &md.ContentType,
		&md.Origin, &md.URL, &md.Timestamp, &md.Size,
	)
	return &md, err
}
