package api

import (
	"database/sql"
	//"github.com/AaronBrownDev/HumaneSociety/internal/repository"
	"net/http"
)

func NewRouter(db *sql.DB) http.Handler {
	// Initialize Repos
	//dogRepo := repository.NewDogRepository(db)
	//personRepo := repository.NewPersonRepository(db)

	// TODO Initialize handlers with their respective repos

	// TODO Implement rest
	return nil
}
