package schema

import (
	"database/sql"
	"fmt"
)

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
