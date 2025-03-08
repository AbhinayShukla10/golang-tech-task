package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AbhinayShukla10/port-service/port-service/internal/domain"
)

// LoadPortsFromJSON loads port data from a JSON file into the repository
func LoadPortsFromJSON(repo *InMemoryPortRepository, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Decode JSON into a map
	decoder := json.NewDecoder(file)
	var ports map[string]domain.Port
	if err := decoder.Decode(&ports); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	// Load ports into the repository
	repo.LoadPorts(ports)
	fmt.Println("✅ Ports loaded successfully.")
	return nil
}
