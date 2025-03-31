package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/AaronBrownDev/HumaneSociety/internal/cmd"
)

func main() {
	// Create a cancellable context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// Execute and get exit code
	exitCode := cmd.Execute(ctx)

	log.Println("Application terminated")
	os.Exit(exitCode)
}
