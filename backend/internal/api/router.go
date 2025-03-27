package api

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/repository"
	"net/http"
)

func NewRouter(db *sql.DB) http.Handler {
	// Initialize Repos
	repoStorage := repository.NewMSSQLStorage(db)

	// TODO Initialize handlers with their respective repos

	// TODO Implement rest
	return nil
}
