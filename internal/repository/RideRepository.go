package repository

import (
	"database/sql"

	"github.com/CarlosGenuino/uber-go/internal/domain"
	"github.com/CarlosGenuino/uber-go/utils"
	"github.com/jmoiron/sqlx"
)

type RideRepository struct {
	db *sqlx.DB
}

func NewRideRepository(db *sqlx.DB) *RideRepository {
	return &RideRepository{db: db}
}

func (r *RideRepository) Save(ride *domain.Ride) error {
	ride.ID = utils.NewUUIDv7()
	query := `INSERT INTO rides (id, passenger_id, driver_id, start_time, end_time, status) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, ride.ID, ride.PassengerID, ride.DriverID, ride.StartTime, ride.EndTime, ride.Status)
	return err
}

func (r *RideRepository) FindByID(id string) (*domain.Ride, error) {
	var ride domain.Ride
	query := `SELECT id, passenger_id, driver_id, start_time, end_time, status FROM rides WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&ride.ID, &ride.PassengerID, &ride.DriverID, &ride.StartTime, &ride.EndTime, &ride.Status)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &ride, err
}

func (r *RideRepository) Update(ride *domain.Ride) error {
	query := `UPDATE rides SET passenger_id = $1, driver_id = $2, start_time = $3, end_time = $4, status = $5 
              WHERE id = $6`
	_, err := r.db.Exec(query, ride.PassengerID, ride.DriverID, ride.StartTime, ride.EndTime, ride.Status, ride.ID)
	return err
}
