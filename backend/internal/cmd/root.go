package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func Execute(ctx context.Context) int {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Using system environment variables.")
	}

	// Create a channel to receive OS signals
	quit := make(chan os.Signal, 1)
	// Add interrupt signals (CTRL + C) or termination signals to channel
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start API server
	go func() {
		if err := APICmd(ctx); err != nil {
			log.Printf("Error starting API: %v", err)
			quit <- syscall.SIGTERM
		}
	}()

	// Wait until signal is received
	<-quit
	log.Println("Shutting down gracefully...")

	// Return success
	return 0
}
