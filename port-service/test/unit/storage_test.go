package infrastructure_test

import (
	"testing"

	"github.com/AbhinayShukla10/port-service/port-service/internal/domain"
	"github.com/AbhinayShukla10/port-service/port-service/internal/infrastructure"
	"github.com/stretchr/testify/assert"
)

// TestNewInMemoryPortRepository ensures the repository initializes correctly
func TestNewInMemoryPortRepository(t *testing.T) {
	repo := infrastructure.NewInMemoryPortRepository()
	assert.NotNil(t, repo)
	// assert.Equal(t, 0, len(repo.GetAllPorts()))
}

// TestLoadPorts checks if ports are loaded correctly
func TestLoadPorts(t *testing.T) {
	repo := infrastructure.NewInMemoryPortRepository()
	ports := map[string]domain.Port{
		"AEAJM": {
			Name:     "Ajman",
			City:     "Ajman",
			Country:  "United Arab Emirates",
			Unlocs:   []string{"AEAJM"},
			Code:     "52000",
			Timezone: "Asia/Dubai",
		},
	}

	repo.LoadPorts(ports)
	port, exists := repo.GetPort("AEAJM")
	assert.True(t, exists)
	assert.Equal(t, "Ajman", port.Name)
}

// TestGetPort checks retrieval of an existing port
func TestGetPort(t *testing.T) {
	repo := infrastructure.NewInMemoryPortRepository()
	repo.LoadPorts(map[string]domain.Port{
		"ZWBUQ": {Name: "Bulawayo", City: "Bulawayo", Country: "Zimbabwe"},
	})

	port, exists := repo.GetPort("ZWBUQ")
	assert.True(t, exists)
	assert.Equal(t, "Bulawayo", port.Name)
}

// TestGetPortNotFound checks retrieval of a non-existing port
func TestGetPortNotFound(t *testing.T) {
	repo := infrastructure.NewInMemoryPortRepository()
	_, exists := repo.GetPort("XYZ")
	assert.False(t, exists)
}

// TestLoadPortsOverwrite checks if loading new ports overwrites old data
func TestLoadPortsOverwrite(t *testing.T) {
	repo := infrastructure.NewInMemoryPortRepository()
	repo.LoadPorts(map[string]domain.Port{"AEAJM": {Name: "Ajman"}})
	repo.LoadPorts(map[string]domain.Port{"ZWHRE": {Name: "Harare"}})

	_, existsOld := repo.GetPort("AEAJM")
	assert.False(t, existsOld)

	port, existsNew := repo.GetPort("ZWHRE")
	assert.True(t, existsNew)
	assert.Equal(t, "Harare", port.Name)
}
