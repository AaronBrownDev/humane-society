package repository

import (
	"context"
	"database/sql"
	"fmt"
)

// withTx executes a function within a transaction
func withTx(db *sql.DB, ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction failed and could not roll back: %w: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
