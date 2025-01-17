package domain

import "fmt"

// Passenger represents a passenger in the system
type Passenger struct {
	ID       string
	Name     string
	Location *Location
}

// NewPassenger creates a new Passenger instance
func NewPassenger(id, name string, latitude, longitude float64) *Passenger {
	return &Passenger{
		ID:       id,
		Name:     name,
		Location: NewLocation(latitude, longitude),
	}
}

// RequestRide simulates a passenger requesting a ride
func (p *Passenger) RequestRide(driver *Driver) {
	if !driver.IsAvailable() {
		fmt.Printf("Passenger %s cannot request a ride. Driver %s is not available.\n", p.Name, driver.Name)
		return
	}

	distance := p.Location.DistanceTo(driver.Location)
	fmt.Printf("Passenger %s is %.2f km away from driver %s\n", p.Name, distance, driver.Name)

	driver.StartRide()
	fmt.Printf("Passenger %s requested a ride with driver %s\n", p.Name, driver.Name)
}

// EndRide simulates a passenger ending a ride
func (p *Passenger) EndRide(driver *Driver) {
	if driver.IsAvailable() {
		fmt.Printf("Passenger %s is not on a ride with driver %s\n", p.Name, driver.Name)
		return
	}

	driver.EndRide()
	fmt.Printf("Passenger %s ended the ride with driver %s\n", p.Name, driver.Name)
}

// UpdateLocation updates the passenger's location
func (p *Passenger) UpdateLocation(latitude, longitude float64) {
	p.Location = NewLocation(latitude, longitude)
	fmt.Printf("Passenger %s location updated to (%f, %f)\n", p.Name, latitude, longitude)
}
