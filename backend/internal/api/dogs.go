package api

import (
	"context"
	"encoding/json"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func (a *api) getAllDogsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	dogs, err := a.repositories.Dogs.GetAll(ctx)
	if err != nil {
		a.logger.Printf("error getting all dogs: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, dogs)
}

func (a *api) getAvailableDogsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	dogs, err := a.repositories.Dogs.GetAvailable(ctx)
	if err != nil {
		a.logger.Printf("error getting available dogs: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, dogs)
}

func (a *api) getDogByIDHandler(w http.ResponseWriter, r *http.Request) {
	dogID, err := uuid.Parse(chi.URLParam(r, "dogID"))
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid Dog ID")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	dog, err := a.repositories.Dogs.GetByID(ctx, dogID)
	if err != nil {
		a.logger.Printf("error getting dog by ID: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
	}

	a.respondJSON(w, http.StatusOK, dog)
}

func (a *api) createDogHandler(w http.ResponseWriter, r *http.Request) {
	var dog domain.Dog
	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.Dogs.Create(ctx, &dog); err != nil {
		a.logger.Printf("error creating dog: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *api) updateDogHandler(w http.ResponseWriter, r *http.Request) {
	dogID, err := uuid.Parse(chi.URLParam(r, "dogID"))
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid Dog ID")
		return
	}

	var dog domain.Dog
	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		a.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	dog.DogID = dogID

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.Dogs.Update(ctx, &dog); err != nil {
		a.logger.Printf("error updating dog: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
	}

	a.respondJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *api) deleteDogHandler(w http.ResponseWriter, r *http.Request) {
	dogID, err := uuid.Parse(chi.URLParam(r, "dogID"))
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid Dog ID")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.Dogs.Delete(ctx, dogID); err != nil {
		a.logger.Printf("error deleting dog: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
	}

	a.respondJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *api) markAdoptedDogHandler(w http.ResponseWriter, r *http.Request) {
	dogID, err := uuid.Parse(chi.URLParam(r, "dogID"))
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "Invalid Dog ID")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.Dogs.MarkAsAdopted(ctx, dogID); err != nil {
		a.logger.Printf("error marking as adopted: %v", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
