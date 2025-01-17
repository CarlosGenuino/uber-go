package domain

type Car struct {
	Make  string
	Model string
	Year  int
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
