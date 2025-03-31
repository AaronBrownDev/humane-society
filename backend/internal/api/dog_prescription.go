package api

import (
	"context"
	"encoding/json"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func (a *api) getAllDogPrescriptionsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	prescriptions, err := a.repositories.DogPrescriptions.GetAll(ctx)
	if err != nil {
		a.logger.Println("error getting all dog prescriptions: ", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, prescriptions)
}

func (a *api) getActiveDogPrescriptionsHandler(w http.ResponseWriter, r *http.Request) {
	dogID, err := uuid.Parse(chi.URLParam(r, "dogID"))
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "invalid dog prescription id")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	prescriptions, err := a.repositories.DogPrescriptions.GetActivePrescriptions(ctx, dogID)
	if err != nil {
		a.logger.Println("error getting active dog prescriptions: ", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	a.respondJSON(w, http.StatusOK, prescriptions)
}

func (a *api) getDogPrescriptionByIDHandler(w http.ResponseWriter, r *http.Request) {
	prescriptionID, err := strconv.Atoi(chi.URLParam(r, "prescriptionID"))
	if err != nil {
		// TODO: look into if logger is needed
		// a.logger.Println("error getting dog prescription id: ", err)
		a.respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	prescription, err := a.repositories.DogPrescriptions.GetByID(ctx, prescriptionID)
	if err != nil {
		a.logger.Println("error getting dog prescription: ", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, prescription)
}

func (a *api) getDogPrescriptionsByDogIDHandler(w http.ResponseWriter, r *http.Request) {
	dogID, err := uuid.Parse(chi.URLParam(r, "dogID"))
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	prescriptions, err := a.repositories.DogPrescriptions.GetByDogID(ctx, dogID)
	if err != nil {
		a.logger.Println("error getting dog prescription: ", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, prescriptions)
}

func (a *api) createDogPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	dogID, err := uuid.Parse(chi.URLParam(r, "dogID"))
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var prescription domain.DogPrescription
	if err := json.NewDecoder(r.Body).Decode(&prescription); err != nil {
		a.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	prescription.DogID = dogID

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.DogPrescriptions.Create(ctx, &prescription); err != nil {
		a.logger.Println("error creating dog prescription: ", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *api) updateDogPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	prescriptionID, err := strconv.Atoi(chi.URLParam(r, "prescriptionID"))
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var prescription domain.DogPrescription
	if err := json.NewDecoder(r.Body).Decode(&prescription); err != nil {
		a.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	prescription.PrescriptionID = prescriptionID

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.DogPrescriptions.Update(ctx, &prescription); err != nil {
		a.logger.Println("error updating dog prescription: ", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *api) deleteDogPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	prescriptionID, err := strconv.Atoi(chi.URLParam(r, "prescriptionID"))
	if err != nil {
		a.respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if err := a.repositories.DogPrescriptions.Delete(ctx, prescriptionID); err != nil {
		a.logger.Println("error deleting dog prescription: ", err)
		a.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.respondJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
