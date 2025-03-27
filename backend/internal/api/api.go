package api

import (
	"database/sql"
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
		// TODO: initialize handlers
	})

	// Dog prescription routes
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
