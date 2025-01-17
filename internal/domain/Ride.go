package domain

import "time"

type RideStatus string

const (
	RideStatusRequested  RideStatus = "requested"
	RideStatusInProgress RideStatus = "in_progress"
	RideStatusCompleted  RideStatus = "completed"
	RideStatusCanceled   RideStatus = "canceled"
)

type Ride struct {
	ID          string
	PassengerID string
	DriverID    string
	StartTime   time.Time
	EndTime     time.Time
	Status      RideStatus
}

func NewRide(id, passengerID, driverID string) *Ride {
	return &Ride{
		ID:          id,
		PassengerID: passengerID,
		DriverID:    driverID,
		StartTime:   time.Now(), // Set the start time to the current time
		Status:      RideStatusRequested,
	}
}

func (r *Ride) Start() {
	r.Status = RideStatusInProgress
	r.StartTime = time.Now()
}

func (r *Ride) End() {
	r.Status = RideStatusCompleted
	r.EndTime = time.Now()
}

func (r *Ride) Cancel() {
	r.Status = RideStatusCanceled
}
