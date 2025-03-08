package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AbhinayShukla10/port-service/port-service/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock repository implementing the PortRepository interface
type mockPortRepository struct {
	ports map[string]domain.Port
}

func (m *mockPortRepository) GetPort(id string) (domain.Port, bool) {
	port, exists := m.ports[id]
	return port, exists
}

func TestStartServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock data
	mockRepo := &mockPortRepository{
		ports: map[string]domain.Port{
			"AEAJM": {
				Name:        "Ajman",
				City:        "Ajman",
				Country:     "United Arab Emirates",
				Province:    "Ajman",
				Timezone:    "Asia/Dubai",
				Unlocs:      []string{"AEAJM"},
				Code:        "52000",
				Coordinates: []float64{55.5136433, 25.4052165},
			},
		},
	}

	// Create a test server
	router := gin.Default()
	router.GET("/ports/:id", func(c *gin.Context) {
		id := c.Param("id")
		port, exists := mockRepo.GetPort(id)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Port not found"})
			return
		}
		c.JSON(http.StatusOK, port)
	})

	t.Run("Fetch existing port", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/ports/AEAJM", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Ajman") // Ensure response contains correct port details
	})

	t.Run("Fetch non-existent port", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/ports/INVALID", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Port not found")
	})
}
