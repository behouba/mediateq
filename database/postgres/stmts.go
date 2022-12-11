package postgres

import "database/sql"

const (
	insertMediaSQL = ""
)

type stmts struct {
	insertMediaStmt *sql.Stmt
	selectMediaStmt *sql.Stmt
	deleteMediaStmt *sql.Stmt
}
