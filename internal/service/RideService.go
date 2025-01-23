package service

import (
	"errors"

	"github.com/CarlosGenuino/uber-go/internal/domain"
	"github.com/CarlosGenuino/uber-go/internal/repository"
)

type RideService struct {
	passengerService *PassengerService
	driverService    *DriverService
	rideRepo         *repository.RideRepository
}

func NewRideService(passengerService *PassengerService, driverService *DriverService, rideRepo *repository.RideRepository) *RideService {
	return &RideService{
		passengerService: passengerService,
		driverService:    driverService,
		rideRepo:         rideRepo,
	}
}

func (s *RideService) RequestRide(passengerID string) (*domain.Ride, error) {
	passenger, err := s.passengerService.GetPassenger(passengerID)
	if err != nil {
		return nil, errors.New("passenger not found")
	}

	ride := &domain.Ride{
		PassengerID: passenger.ID,
		Status:      domain.RideStatusRequested,
	}
	if err := s.rideRepo.Save(ride); err != nil {
		return nil, err
	}
	return ride, nil
}

func (s *RideService) AcceptRide(rideID, driverID string) (*domain.Ride, error) {
	ride, err := s.rideRepo.FindByID(rideID)
	if err != nil {
		return nil, errors.New("ride not found")
	}

	driver, err := s.driverService.GetDriver(driverID)
	if err != nil {
		return nil, errors.New("driver not found")
	}

	if !driver.Available {
		return nil, errors.New("driver is not available")
	}

	ride.DriverID = driver.ID
	ride.Status = domain.RideStatusInProgress
	driver.Available = false

	if err := s.rideRepo.Update(ride); err != nil {
		return nil, err
	}
	return ride, nil
}

func (s *RideService) EndRide(rideID string) (*domain.Ride, error) {
	ride, err := s.rideRepo.FindByID(rideID)
	if err != nil {
		return nil, errors.New("ride not found")
	}

	driver, err := s.driverService.GetDriver(ride.DriverID)
	if err != nil {
		return nil, errors.New("driver not found")
	}

	ride.Status = domain.RideStatusCompleted
	driver.Available = true

	if err := s.rideRepo.Update(ride); err != nil {
		return nil, err
	}
	return ride, nil
}

func (s *RideService) CancelRide(rideID string) (*domain.Ride, error) {
	ride, err := s.rideRepo.FindByID(rideID)
	if err != nil {
		return nil, errors.New("ride not found")
	}

	if ride.DriverID != "" {
		driver, err := s.driverService.GetDriver(ride.DriverID)
		if err != nil {
			return nil, errors.New("driver not found")
		}
		driver.Available = true
	}

	ride.Status = domain.RideStatusCanceled

	if err := s.rideRepo.Update(ride); err != nil {
		return nil, err
	}
	return ride, nil
}
