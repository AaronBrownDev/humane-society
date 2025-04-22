package api

import (
	"context"
	"net/http"
	"time"
)

// getAllUserAccountsHandler returns a list of all user accounts in the system
func (a *api) getAllUserAccountsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	// Get all user accounts from the repository
	userAccounts, err := a.repositories.UserAccounts.GetAll(ctx)
	if err != nil {
		a.logger.Printf("Error getting all user accounts: %v", err)
		a.respondError(w, http.StatusInternalServerError, "Failed to retrieve user accounts")
		return
	}

	// Return the user accounts as JSON
	a.respondJSON(w, http.StatusOK, userAccounts)
}
