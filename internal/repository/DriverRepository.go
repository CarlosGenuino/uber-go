package repository

import (
	"errors"

	"github.com/CarlosGenuino/uber-go/internal/domain"
)

type DriverRepository struct {
	drivers map[string]*domain.Driver
}

func NewDriverRepository() *DriverRepository {
	return &DriverRepository{
		drivers: make(map[string]*domain.Driver),
	}
}

func (r *DriverRepository) Save(driver *domain.Driver) error {
	if _, exists := r.drivers[driver.ID]; exists {
		return errors.New("driver already exists")
	}
	r.drivers[driver.ID] = driver
	return nil
}

func (r *DriverRepository) FindByID(id string) (*domain.Driver, error) {
	driver, exists := r.drivers[id]
	if !exists {
		return nil, errors.New("driver not found")
	}
	return driver, nil
}
