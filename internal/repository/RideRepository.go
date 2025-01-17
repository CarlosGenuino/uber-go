package repository

import (
	"errors"

	"github.com/CarlosGenuino/uber-go/internal/domain"
)

type RideRepository struct {
	rides map[string]*domain.Ride
}

func NewRideRepository() *RideRepository {
	return &RideRepository{
		rides: make(map[string]*domain.Ride),
	}
}

func (r *RideRepository) Save(ride *domain.Ride) error {
	if _, exists := r.rides[ride.ID]; exists {
		return errors.New("ride already exists")
	}
	r.rides[ride.ID] = ride
	return nil
}

func (r *RideRepository) FindByID(id string) (*domain.Ride, error) {
	ride, exists := r.rides[id]
	if !exists {
		return nil, errors.New("ride not found")
	}
	return ride, nil
}

func (r *RideRepository) Update(ride *domain.Ride) error {
	if _, exists := r.rides[ride.ID]; !exists {
		return errors.New("ride not found")
	}
	r.rides[ride.ID] = ride
	return nil
}
