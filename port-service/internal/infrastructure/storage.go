package infrastructure

import (
	"sync"

	"github.com/AbhinayShukla10/port-service/port-service/internal/domain"
)

// InMemoryPortRepository is an in-memory implementation of PortRepository
type InMemoryPortRepository struct {
	mu    sync.RWMutex
	ports map[string]domain.Port
}

// NewInMemoryPortRepository initializes an empty repository
func NewInMemoryPortRepository() *InMemoryPortRepository {
	return &InMemoryPortRepository{
		ports: make(map[string]domain.Port),
	}
}

// GetPort retrieves a port by its ID
func (r *InMemoryPortRepository) GetPort(id string) (domain.Port, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	port, exists := r.ports[id]
	return port, exists
}

// LoadPorts loads port data from a JSON map
func (r *InMemoryPortRepository) LoadPorts(data map[string]domain.Port) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ports = data
}
