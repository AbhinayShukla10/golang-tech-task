package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbhinayShukla10/port-service/port-service/internal/domain"
)

// StartServer initializes the HTTP server using only the standard library
func StartServer(repo domain.PortRepository) {
	http.HandleFunc("/ports/", func(w http.ResponseWriter, r *http.Request) {
		// Ensure it's a GET request
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract port ID from URL
		id := r.URL.Path[len("/ports/"):]
		if id == "" {
			http.Error(w, "Port ID is required", http.StatusBadRequest)
			return
		}

		// Retrieve the port from the repository
		port, exists := repo.GetPort(id)
		if !exists {
			http.Error(w, "Port not found", http.StatusNotFound)
			return
		}

		// Convert struct to JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(port); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	// Start server on port 8080
	fmt.Println("🚀 Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
