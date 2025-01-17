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

func (s *RideService) RequestRide(passengerID, driverID string) (*domain.Ride, error) {
	_, err := s.passengerService.GetPassenger(passengerID)
	if err != nil {
		return nil, errors.New("passenger not found")
	}

	driver, err := s.driverService.GetDriver(driverID)
	if err != nil {
		return nil, errors.New("driver not found")
	}

	if !driver.IsAvailable() {
		return nil, errors.New("driver is not available")
	}

	// Create a new ride
	ride := domain.NewRide(generateID(), passengerID, driverID)

	// Start the ride
	ride.Start()
	driver.StartRide()

	// Save the ride
	if err := s.rideRepo.Save(ride); err != nil {
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

	// End the ride
	ride.End()
	driver.EndRide()

	// Update the ride in the repository
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

	driver, err := s.driverService.GetDriver(ride.DriverID)
	if err != nil {
		return nil, errors.New("driver not found")
	}

	// Cancel the ride
	ride.Cancel()
	driver.EndRide() // Mark the driver as available again

	// Update the ride in the repository
	if err := s.rideRepo.Update(ride); err != nil {
		return nil, err
	}

	return ride, nil
}

func generateID() string {
	// Implement ID generation logic (e.g., UUID)
	return "ride-id"
}
