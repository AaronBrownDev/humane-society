package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"time"
)

type api struct {
	repositories *repository.Storage
	db           *sql.DB
	logger       *log.Logger
}

func NewAPI(db *sql.DB) *api {
	return &api{
		repositories: repository.NewMSSQLStorage(db),
		db:           db,
		logger:       log.New(os.Stdout, "api: ", log.LstdFlags),
	}
}

func (a *api) Routes() http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	// Health check
	router.Get("/health", a.healthCheckHandler)

	// Dog routes
	router.Route("/api/dogs", func(r chi.Router) {
		r.Get("/", a.getAllDogsHandler)
		r.Get("/available", a.getAvailableDogsHandler)
		r.Get("/{dogID}", a.getDogByIDHandler)

		r.Post("/", a.createDogHandler)
		r.Put("/{dogID}", a.updateDogHandler)
		r.Delete("/{dogID}", a.deleteDogHandler)
		r.Patch("/{dogID}/adopt", a.markAdoptedDogHandler)

		// TODO implement prescription handlers
		// DogID prescription routes
		r.Route("/{dogID}/prescriptions", func(r chi.Router) {
			r.Get("/", a.getDogPrescriptionsByDogIDHandler)
			r.Get("/active", a.getActiveDogPrescriptionsHandler)

			r.Post("/", a.createDogPrescriptionHandler)
		})
	})

	// Dog prescription routes
	router.Route("/api/prescriptions", func(r chi.Router) {
		r.Get("/", a.getAllDogPrescriptionsHandler)
		r.Get("/{prescriptionID}", a.getDogPrescriptionByIDHandler)

		r.Put("/{prescriptionID}", a.updateDogPrescriptionHandler)
		r.Delete("/{prescriptionID}", a.deleteDogPrescriptionHandler)
	})

	// TODO: finish the routes
	router.Route("/api/", func(r chi.Router) {})

	return router
}

func (a *api) Server(port int) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      a.Routes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  time.Minute,
	}
}

func (a *api) respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			a.logger.Printf("error encoding response: %v", err)
		}
	}
}

func (a *api) respondError(w http.ResponseWriter, code int, message string) {
	a.respondJSON(w, code, map[string]string{"error": message})
}
