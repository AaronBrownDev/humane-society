package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/repository"
	"github.com/AaronBrownDev/HumaneSociety/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type api struct {
	repositories *repository.Storage
	authService  *services.AuthService
	db           *sql.DB
	logger       *log.Logger
}

func NewAPI(db *sql.DB) *api {
	repos := repository.NewMSSQLStorage(db)

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		log.Fatalf("Set JWT_SECRET environment variable")
	}

	return &api{
		repositories: repos,
		authService: services.NewAuthService(
			repos.People,
			repos.UserAccounts,
			repos.RefreshTokens,
			repos.Roles,
			repos.UserRoles,
			jwtSecret,
			15,
			7,
		),
		db:     db,
		logger: log.New(os.Stdout, "api: ", log.LstdFlags),
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
	router.Use(corsMiddleware)
	// Health check
	router.Get("/health", a.healthCheckHandler)

	// Auth routes
	router.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", a.registerHandler)
		r.Post("/login", a.loginHandler)
		r.Post("/refresh", a.refreshTokenHandler)
		r.Post("/logout", a.logoutHandler)
	})

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
	// User account routes
	router.Route("/api/users", func(r chi.Router) {
		r.Get("/", a.getAllUserAccountsHandler)
	})

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

// corsMiddleware handles Cross-Origin Resource Sharing (CORS) for the API
// It correctly configures headers for secure cross-origin requests with credentials
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get allowed origins from environment or use default
		allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
		if allowedOrigins == "" {
			allowedOrigins = "http://localhost:5173" // Default for local development
		}

		// Check if the request origin is allowed
		origin := r.Header.Get("Origin")
		if origin != "" {
			// Check if origin matches any allowed origin
			originAllowed := false
			for _, allowedOrigin := range strings.Split(allowedOrigins, ",") {
				if origin == strings.TrimSpace(allowedOrigin) {
					originAllowed = true
					break
				}
			}

			if originAllowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			} else {
				// If origin is not allowed, default to the first allowed origin
				// This is not ideal but better than "*" which would prevent credentials
				w.Header().Set("Access-Control-Allow-Origin", strings.TrimSpace(strings.Split(allowedOrigins, ",")[0]))
			}
		} else {
			// If no origin header, use the first allowed origin
			w.Header().Set("Access-Control-Allow-Origin", strings.TrimSpace(strings.Split(allowedOrigins, ",")[0]))
		}

		// Essential CORS headers
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
