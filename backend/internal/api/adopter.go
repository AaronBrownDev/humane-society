package api

import (
	"context"
	"encoding/json"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func (a *api) getAllAdoptersHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	adopters, err := a.repositories.Adopters.GetAll(ctx)
	if err != nil {
		a.logger.Printf("error getting all adopters: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, adopters)
}

func (a *api) getAdopterByIDHandler(w http.ResponseWriter, r *http.Request) {
	adopterID, err := uuid.Parse(chi.URLParam(r, "adopterID"))
	if err != nil {
		a.logger.Printf("error parsing adopterID: %v", err)
		a.respondError(w, http.StatusBadRequest, "Invalid Adopter ID")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	adopter, err := a.repositories.Adopters.GetByID(ctx, adopterID)
	if err != nil {
		a.logger.Printf("error getting adopter: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, adopter)
}

func (a *api) createAdopterHandler(w http.ResponseWriter, r *http.Request) {
	var adopter domain.Adopter
	if err := json.NewDecoder(r.Body).Decode(&adopter); err != nil {
		a.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.Adopters.Create(ctx, &adopter); err != nil {
		a.logger.Printf("error creating adopter: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *api) updateAdopterHandler(w http.ResponseWriter, r *http.Request) {
	adopterID, err := uuid.Parse(chi.URLParam(r, "adopterID"))
	if err != nil {
		a.logger.Printf("error parsing adopterID: %v", err)
		a.respondError(w, http.StatusBadRequest, "Invalid Adopter ID")
		return
	}

	var adopter domain.Adopter
	if err := json.NewDecoder(r.Body).Decode(&adopter); err != nil {
		a.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	adopter.AdopterID = adopterID

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.Adopters.Update(ctx, &adopter); err != nil {
		a.logger.Printf("error updating adopter: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *api) deleteAdopterHandler(w http.ResponseWriter, r *http.Request) {
	adopterID, err := uuid.Parse(chi.URLParam(r, "adopterID"))
	if err != nil {
		a.logger.Printf("error parsing adopterID: %v", err)
		a.respondError(w, http.StatusBadRequest, "Invalid Adopter ID")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.Adopters.Delete(ctx, adopterID); err != nil {
		a.logger.Printf("error deleting adopter: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
