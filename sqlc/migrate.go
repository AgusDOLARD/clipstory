package storage

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed schema.sql
var ddl string

func Migrate(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}
	return nil
}
