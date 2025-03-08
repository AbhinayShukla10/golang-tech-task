package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/AbhinayShukla10/port-service/port-service/internal/domain"
	"github.com/gorilla/mux"
)

// PortHandler manages port-related HTTP endpoints.
type PortHandler struct {
	Service *domain.PortService
}

// NewPortHandler creates a new PortHandler.
func NewPortHandler(service *domain.PortService) *PortHandler {
	return &PortHandler{Service: service}
}

// GetPort handles GET /ports/{id} requests.
func (h *PortHandler) GetPort(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	port, found := h.Service.GetPort(id)

	if !found {
		http.Error(w, "Port not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(port)
}
