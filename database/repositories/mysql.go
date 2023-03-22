package repositories

import (
	"context"
	"database/sql"
)

func Prepare(tx *sql.Tx, conn *sql.Conn, query string) (stmt *sql.Stmt, err error) {
	ctx := context.TODO()

	if conn == nil {
		stmt, err = tx.PrepareContext(ctx, query)
	} else {
		stmt, err = conn.PrepareContext(ctx, query)
	}

	return
}
