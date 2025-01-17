package repository

import (
	"errors"

	"github.com/CarlosGenuino/uber-go/internal/domain"
)

type PassengerRepository struct {
	passengers map[string]*domain.Passenger
}

func NewPassengerRepository() *PassengerRepository {
	return &PassengerRepository{
		passengers: make(map[string]*domain.Passenger),
	}
}

func (r *PassengerRepository) Save(passenger *domain.Passenger) error {
	if _, exists := r.passengers[passenger.ID]; exists {
		return errors.New("passenger already exists")
	}
	r.passengers[passenger.ID] = passenger
	return nil
}

func (r *PassengerRepository) FindByID(id string) (*domain.Passenger, error) {
	passenger, exists := r.passengers[id]
	if !exists {
		return nil, errors.New("passenger not found")
	}
	return passenger, nil
}
