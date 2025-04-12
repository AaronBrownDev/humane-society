package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/AaronBrownDev/HumaneSociety/internal/api"
	"github.com/AaronBrownDev/HumaneSociety/internal/database"
)

func APICmd(ctx context.Context) error {
	// Ensure database exists
	if err := database.EnsureDatabaseExists(); err != nil {
		return fmt.Errorf("error ensuring database exists: %v", err)
	}

	// Initialize database
	if err := database.InitializeDB(); err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}
	defer database.CloseDB()

	// Run migrations
	if err := database.RunMigrations(); err != nil {
		return fmt.Errorf("error running migrations: %v", err)
	}

	// Get database
	db := database.GetDB()

	// Defines the port number for the server from environment or uses default
	port := 8080
	if os.Getenv("PORT") != "" {
		var err error
		port, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Printf("Invalid PORT value, using default: %d", port)
		}
	}

	// Create API server
	apiServer := api.NewAPI(db)
	server := apiServer.Server(port)

	log.Printf("Starting server on port %d", port)

	// Starts server in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	// Wait for termination signal or server error
	select {
	case <-ctx.Done():
		log.Println("Shutdown signal received")

		// Gives time for connections to end
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Stops accepting new connections
		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server forced to shutdown: %v", err)
		}

		log.Println("Server shut down")
		return nil

	case err := <-serverErrors:
		return fmt.Errorf("error starting server: %v", err)
	}
}
