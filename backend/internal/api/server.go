package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server represents the HTTP server and its dependencies
type Server struct {
	router     http.Handler // HTTP server router/mux
	httpServer *http.Server // HTTP server instance
	db         *sql.DB      // database connection pool
}

// NewServer returns a new Server instance
func NewServer(db *sql.DB) *Server {
	router := NewRouter(db)

	return &Server{
		router: router,
		db:     db,
	}
}

// Run starts the HTTP server and waits for shutdown
// returns error if shutdown encounters an error
func (s *Server) Run() error {
	// Defines the port number for the server from environment or uses default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Defines the server params
	s.httpServer = &http.Server{
		Addr:         ":" + port,       // Port address
		Handler:      s.router,         // Router
		ReadTimeout:  15 * time.Second, // Max duration for read requests
		WriteTimeout: 15 * time.Second, // Max duration for write responses
		IdleTimeout:  time.Minute,      // Max duration to wait for next request
	}

	// Starts server
	go func() {
		log.Printf("Starting server on port %s", port)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Channel that receives signal for shutdown
	quit := make(chan os.Signal, 1)
	// Adds interrupt signals (CTRL + C) or termination signals to channel
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Waits until signal is received
	<-quit
	log.Println("Shutting down server...")

	// Gives 15 seconds for connections to end
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Stops accepting new connections
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %v", err)
	}

	log.Println("Server shut down")
	return nil
}
