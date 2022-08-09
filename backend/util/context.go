package util

import (
	"context"
	"database/sql"
)

// GetDBTx :
func GetDBTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value("dbTx").(*sql.Tx); ok {
		return tx
	}
	return nil
}
