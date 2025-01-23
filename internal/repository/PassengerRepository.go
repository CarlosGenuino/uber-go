package repository

import (
	"database/sql"

	"github.com/CarlosGenuino/uber-go/internal/domain"
	"github.com/CarlosGenuino/uber-go/utils"
	"github.com/jmoiron/sqlx"
)

type PassengerRepository struct {
	db *sqlx.DB
}

func NewPassengerRepository(db *sqlx.DB) *PassengerRepository {
	return &PassengerRepository{db: db}
}

func (r *PassengerRepository) Save(passenger *domain.Passenger) error {
	passenger.ID = utils.NewUUIDv7()
	query := `INSERT INTO passengers (id, name, latitude, longitude) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, passenger.ID, passenger.Name, passenger.Location.Latitude, passenger.Location.Longitude)
	return err
}

func (r *PassengerRepository) FindByID(id string) (*domain.Passenger, error) {
	var passenger domain.Passenger
	var lat, lon float64
	query := `SELECT id, name, latitude, longitude FROM passengers WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&passenger.ID, &passenger.Name, &lat, &lon)
	if err == sql.ErrNoRows {
		return nil, err
	}
	passenger.Location = domain.NewLocation(lat, lon)
	return &passenger, err
}
