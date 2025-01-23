package service

import (
	"github.com/CarlosGenuino/uber-go/internal/domain"
	"github.com/CarlosGenuino/uber-go/internal/repository"
)

type PassengerService struct {
	repo *repository.PassengerRepository
}

func NewPassengerService(repo *repository.PassengerRepository) *PassengerService {
	return &PassengerService{repo: repo}
}

func (s *PassengerService) CreatePassenger(name string, latitude, longitude float64) (*domain.Passenger, error) {
	passenger := &domain.Passenger{
		Name:     name,
		Location: domain.NewLocation(latitude, longitude),
	}
	if err := s.repo.Save(passenger); err != nil {
		return nil, err
	}
	return passenger, nil
}

func (s *PassengerService) GetPassenger(id string) (*domain.Passenger, error) {
	return s.repo.FindByID(id)
}
