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
		media_repository (media_id, content_type, origin, url, base64hash, timestamp, size_bytes) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7);
	`

	selectMediaFormat = "SELECT media_id, content_type, origin, url, base64hash, timestamp, size_bytes FROM media_repository %s;"

	deleteMediaSQL = ""
)

var (
	selectMediaListSQL = fmt.Sprintf(selectMediaFormat, "OFFSET $1 LIMIT CASE WHEN $2 = 0 THEN 10 ELSE $2 END;")

	selectMediaByHashSQL = fmt.Sprintf(selectMediaFormat, "WHERE base64hash=$1")

	selectMediaByIDSQL = fmt.Sprintf(selectMediaFormat, "WHERE media_id=$1")
)

type mediaStmts struct {
	insertStmt       *sql.Stmt
	selectByHashStmt *sql.Stmt
	selectByIDStmt   *sql.Stmt
	selectListStmt   *sql.Stmt
	deleteStmt       *sql.Stmt
}

func newMediaTable(db *sql.DB) (schema.MediaTable, error) {
	s := &mediaStmts{}
	return s, schema.StatementList{
		{&s.insertStmt, insertMediaSQL},
		{&s.selectByHashStmt, selectMediaByHashSQL},
		{&s.selectByIDStmt, selectMediaByIDSQL},
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
			&m.ID, &m.Base64Hash, "", &m.ContentType,
			&m.Origin, &m.URL, &m.Timestamp, &m.SizeBytes,
		); err != nil {
			return nil, err
		}
		mediaList = append(mediaList, m)
	}

	return mediaList, nil
}

// Insert implements schema.MediaTable
func (s mediaStmts) Insert(ctx context.Context, m *mediateq.Media) error {
	_, err := s.insertStmt.ExecContext(
		ctx, m.ID, m.ContentType, m.Origin, m.URL, m.Base64Hash, m.Timestamp, m.SizeBytes,
	)
	if err != nil {
		return err
	}
	return nil
}

// SelectByUID implements schema.MediaTable
func (s mediaStmts) SelectByID(ctx context.Context, hash string) (*mediateq.Media, error) {
	md := mediateq.Media{}
	// media_id, content_type, origin, url, base64hash, timestamp, size_bytes
	err := s.selectByIDStmt.QueryRowContext(ctx, hash).Scan(
		&md.ID, &md.ContentType, &md.Origin, &md.URL,
		&md.Base64Hash, &md.Timestamp, &md.SizeBytes,
	)
	return &md, err
}

// SelectByUID implements schema.MediaTable
func (s mediaStmts) SelectByBase64Hash(ctx context.Context, base64Hash string) (*mediateq.Media, error) {
	md := mediateq.Media{}
	err := s.selectByHashStmt.QueryRowContext(ctx, base64Hash).Scan(
		&md.ID, &md.ContentType, &md.Origin, &md.URL,
		&md.Base64Hash, &md.Timestamp, &md.SizeBytes,
	)
	return &md, err
}
