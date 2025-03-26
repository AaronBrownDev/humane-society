package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Dogs       DogRepository
	People     PersonRepository
	Adoptions  AdoptionRepository
	Surrenders SurrenderRepository
	Volunteers VolunteerRepository
	Inventory  InventoryRepository
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Dogs:       NewDogRepository(db),
		People:     NewPersonRepository(db),
		Adoptions:  NewSQLAdoptionRepository(db),
		Surrenders: NewSQLSurrenderRepository(db),
		Volunteers: NewSQLVolunteerRepository(db),
		Inventory:  NewSQLInventoryRepository(db),
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
