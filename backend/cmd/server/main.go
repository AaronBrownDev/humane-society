package main

import (
	"github.com/AaronBrownDev/HumaneSociety/internal/api"
	"github.com/AaronBrownDev/HumaneSociety/internal/database"
	"log"
)

func main() {
	// Initialize database
	if err := database.InitializeDB(); err != nil {
		log.Panicf("Error initializing database: %v", err)
	}
	defer database.CloseDB()

	// Get database
	db := database.GetDB()

	// Start server
	server := api.NewServer(db)
	if err := server.Run(); err != nil {
		log.Panicf("Error starting server: %v", err)
	}
}
