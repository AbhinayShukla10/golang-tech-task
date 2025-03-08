package domain

// PortService handles port-related business logic.
type PortService struct {
	repo PortRepository
}

// NewPortService creates a new PortService.
func NewPortService(repo PortRepository) *PortService {
	return &PortService{repo: repo}
}

// // SavePort saves or updates a port.
// func (s *PortService) SavePort(port Port) error {
// 	return s.repo.SavePort(port)
// }

// GetPort retrieves a port by ID.
func (s *PortService) GetPort(id string) (Port, bool) {
	return s.repo.GetPort(id)
}
