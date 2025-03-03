package main

import (
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/database"
	"log"
)

func main() {
	if err := database.InitializeDB(); err != nil {
		log.Panicf("Error initializing database: %v", err)
	}
	defer database.CloseDB()

	fmt.Println(database.GetDB())

}
