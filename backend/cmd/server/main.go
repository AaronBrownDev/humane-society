package main

import (
	"github.com/AaronBrownDev/HumaneSociety/internal/database"
	"github.com/AaronBrownDev/HumaneSociety/internal/repository"
	"log"
)

func main() {
	if err := database.InitializeDB(); err != nil {
		log.Panicf("Error initializing database: %v", err)
	}
	defer database.CloseDB()

	dogRepo := repository.NewDogRepository(database.GetDB())
}
