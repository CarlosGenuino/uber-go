package repository

import (
	"database/sql"

	"github.com/CarlosGenuino/uber-go/internal/domain"
	"github.com/CarlosGenuino/uber-go/utils"

	"github.com/jmoiron/sqlx"
)

type DriverRepository struct {
	db *sqlx.DB
}

func NewDriverRepository(db *sqlx.DB) *DriverRepository {
	return &DriverRepository{db: db}
}

func (r *DriverRepository) Save(driver *domain.Driver) error {
	driver.ID = utils.NewUUIDv7()
	query := `INSERT INTO drivers (id, name, license_id, available, latitude, longitude, car_make, car_model, car_year) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.Exec(query, driver.ID, driver.Name, driver.LicenseID, driver.Available,
		driver.Location.Latitude, driver.Location.Longitude, driver.Car.Make, driver.Car.Model, driver.Car.Year)
	return err
}

func (r *DriverRepository) FindByID(id string) (*domain.Driver, error) {
	var driver domain.Driver
	var lat, lon float64
	query := `SELECT id, name, license_id, available, latitude, longitude, car_make, car_model, car_year 
              FROM drivers WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&driver.ID, &driver.Name, &driver.LicenseID, &driver.Available,
		&lat, &lon, &driver.Car.Make, &driver.Car.Model, &driver.Car.Year)
	if err == sql.ErrNoRows {
		return nil, err
	}
	driver.Location = domain.NewLocation(lat, lon)
	return &driver, err
}
