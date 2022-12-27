package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/database/schema"
)

const (
	insertThumbnailSQL = `
	INSERT INTO 
		thumbnail (media_id, content_type, origin, url, base64hash, timestamp, size_bytes, width, height, crop) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`

	selectThumbnailFormat = "SELECT media_id, content_type, origin, url, base64hash, timestamp, size_bytes, width, height, crop FROM thumbnail %s;"

	// deleteThumbnailSQL = ""
)

var (
	selectThumbnailSQL = fmt.Sprintf(selectThumbnailFormat, "WHERE media_id=$1 AND width=$2 AND height=$3 AND crop=$4")

	// selectThumbnailByMediaIDSQL = fmt.Sprintf(selectThumbnailFormat, "WHERE media_id=$1")
)

type thumbnailStmts struct {
	insertStmt          *sql.Stmt
	selectStmt          *sql.Stmt
	selectByMediaIDStmt *sql.Stmt
	deleteStmt          *sql.Stmt
}

func newThumbnailTable(db *sql.DB) (schema.ThumbnailTable, error) {
	s := &thumbnailStmts{}
	return s, schema.StatementList{
		{&s.insertStmt, insertThumbnailSQL},
		{&s.selectStmt, selectThumbnailSQL},
	}.Prepare(db)
}

// Delete implements schema.ThumbnailTable
func (*thumbnailStmts) Delete(ctx context.Context, mediaID string) error {
	panic("unimplemented")
}

// Insert implements schema.ThumbnailTable
func (s *thumbnailStmts) Insert(ctx context.Context, tb *mediateq.Thumbnail) error {
	_, err := s.insertStmt.ExecContext(
		ctx, tb.ID, tb.ContentType, tb.Origin, tb.URL, tb.Base64Hash, tb.Timestamp, tb.SizeBytes, tb.Width, tb.Height, tb.Crop,
	)
	return err
}

// Select implements schema.ThumbnailTable
func (s *thumbnailStmts) Select(ctx context.Context, mediaID string, width int, height int, crop bool) (*mediateq.Thumbnail, error) {
	tb := mediateq.Thumbnail{}
	err := s.selectStmt.QueryRowContext(
		ctx, mediaID, width, height, crop,
	).Scan(
		&tb.ID, &tb.ContentType, &tb.Origin, &tb.URL, &tb.Base64Hash, &tb.Timestamp, &tb.SizeBytes, &tb.Width, &tb.Height, &tb.Crop,
	)
	if err != nil {
		return nil, err
	}

	return &tb, nil

}

// SelectByMediaID implements schema.ThumbnailTable
func (*thumbnailStmts) SelectByMediaID(ctx context.Context, mediaID string) ([]mediateq.Thumbnail, error) {
	panic("unimplemented")
}
