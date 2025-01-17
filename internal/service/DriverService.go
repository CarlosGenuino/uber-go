package service

import (
	"github.com/CarlosGenuino/uber-go/internal/domain"
	"github.com/CarlosGenuino/uber-go/internal/repository"
)

type DriverService struct {
	repo *repository.DriverRepository
}

func NewDriverService(repo *repository.DriverRepository) *DriverService {
	return &DriverService{repo: repo}
}

func (s *DriverService) CreateDriver(id, name, licenseID string, latitude, longitude float64) (*domain.Driver, error) {
	driver := domain.NewDriver(id, name, licenseID, latitude, longitude)
	if err := s.repo.Save(driver); err != nil {
		return nil, err
	}
	return driver, nil
}

func (s *DriverService) GetDriver(id string) (*domain.Driver, error) {
	return s.repo.FindByID(id)
}
