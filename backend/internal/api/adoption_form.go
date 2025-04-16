package api

import (
	"context"
	"net/http"
)

// Handler to get all adoption forms
func (a *api) getAllAdoptionFormsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	forms, err := a.repositories.Adoptions.GetAll(ctx)
	if err != nil {
		a.logger.Printf("error getting all adoption forms: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, forms)
}
