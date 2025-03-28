package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
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

// SwapUUIDFormat converts between MSSQL GUID format and standard RFC 4122 UUID format.
// The same function works in both directions because the byte swapping is symmetric.
func SwapUUIDFormat(inputUUID uuid.UUID) (uuid.UUID, error) {
	bytes := make([]byte, 16)
	copy(bytes, inputUUID[:])

	// Swap the first 4 bytes
	bytes[0], bytes[1], bytes[2], bytes[3] = bytes[3], bytes[2], bytes[1], bytes[0]

	// Swap the next 2 bytes
	bytes[4], bytes[5] = bytes[5], bytes[4]

	// Swap the next 2 bytes
	bytes[6], bytes[7] = bytes[7], bytes[6]

	return uuid.FromBytes(bytes)
}
