package integration

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/AbhinayShukla10/port-service/port-service/internal/domain"
	"github.com/AbhinayShukla10/port-service/port-service/internal/infrastructure"
)

// Sample JSON data
const testJSON = `{
	"AEAJM": {
		"name": "Ajman",
		"city": "Ajman",
		"country": "United Arab Emirates",
		"alias": [],
		"regions": [],
		"coordinates": [55.5136433, 25.4052165],
		"province": "Ajman",
		"timezone": "Asia/Dubai",
		"unlocs": ["AEAJM"],
		"code": "52000"
	},
	"USNYC": {
		"name": "New York",
		"city": "New York",
		"country": "United States",
		"alias": [],
		"regions": [],
		"coordinates": [-74.006, 40.7128],
		"province": "New York",
		"timezone": "America/New_York",
		"unlocs": ["USNYC"],
		"code": "10001"
	}
}`

func TestLoadPortsFromFile(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "ports_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Cleanup

	// Write test JSON data
	if _, err := tempFile.WriteString(testJSON); err != nil {
		t.Fatalf("Failed to write JSON: %v", err)
	}
	tempFile.Close()

	// Initialize in-memory repository
	repo := infrastructure.NewInMemoryPortRepository()

	// Read JSON file and unmarshal
	fileData, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var ports map[string]domain.Port
	if err := json.Unmarshal(fileData, &ports); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	// Load ports into repository
	repo.LoadPorts(ports)

	// Expected Ports
	expectedPorts := map[string]domain.Port{
		"AEAJM": {
			Name:        "Ajman",
			City:        "Ajman",
			Country:     "United Arab Emirates",
			Alias:       []string{},
			Regions:     []string{},
			Coordinates: []float64{55.5136433, 25.4052165},
			Province:    "Ajman",
			Timezone:    "Asia/Dubai",
			Unlocs:      []string{"AEAJM"},
			Code:        "52000",
		},
		"USNYC": {
			Name:        "New York",
			City:        "New York",
			Country:     "United States",
			Alias:       []string{},
			Regions:     []string{},
			Coordinates: []float64{-74.006, 40.7128},
			Province:    "New York",
			Timezone:    "America/New_York",
			Unlocs:      []string{"USNYC"},
			Code:        "10001",
		},
	}

	// Validate loaded ports
	for id, expectedPort := range expectedPorts {
		port, exists := repo.GetPort(id)
		if !exists {
			t.Errorf("Port with ID %s not found", id)
			continue
		}

		expectedJSON, _ := json.Marshal(expectedPort)
		actualJSON, _ := json.Marshal(port)
		if string(expectedJSON) != string(actualJSON) {
			t.Errorf("Mismatch for port %s: expected %s, got %s", id, expectedJSON, actualJSON)
		}
	}
}
