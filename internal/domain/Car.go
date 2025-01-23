package domain

type Car struct {
	Make  string `db:"car_make"`
	Model string `db:"car_model"`
	Year  int    `db:"car_year"`
}

func NewCar(make, model string, year int) Car {
	return Car{
		Make:  make,
		Model: model,
		Year:  year,
	}
}

func ValidateCar(car Car) bool {
	if car.Make == "" || car.Model == "" || car.Year == 0 {
		return false
	}
	return true
}
