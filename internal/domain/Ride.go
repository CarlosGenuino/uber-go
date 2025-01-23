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
	ID          string     `db:"id"`
	PassengerID string     `db:"passenger_id"`
	DriverID    string     `db:"driver_id"`
	StartTime   time.Time  `db:"start_time"`
	EndTime     time.Time  `db:"end_time"`
	Status      RideStatus `db:"status"`
}
