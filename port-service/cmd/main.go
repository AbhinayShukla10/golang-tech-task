package main

import (
	"fmt"
	"log"

	"github.com/AbhinayShukla10/port-service/port-service/internal/infrastructure"
)

func main() {
	fmt.Println("📥 Loading ports from JSON...")

	// Initialize in-memory repository
	repo := infrastructure.NewInMemoryPortRepository()

	// Load data from JSON file
	err := infrastructure.LoadPortsFromJSON(repo, "ports.json")
	if err != nil {
		log.Fatalf("Error loading ports: %v", err)
	}

	// Start the server
	infrastructure.StartServer(repo)
}
