package domain

import "fmt"

// Driver represents a driver in the system
type Driver struct {
	ID        string
	Name      string
	LicenseID string
	Available bool
	Location  *Location
}

// NewDriver creates a new Driver instance
func NewDriver(id, name, licenseID string, latitude, longitude float64) *Driver {
	return &Driver{
		ID:        id,
		Name:      name,
		LicenseID: licenseID,
		Available: true,
		Location:  NewLocation(latitude, longitude),
	}
}

// StartRide marks the driver as unavailable
func (d *Driver) StartRide() {
	if !d.Available {
		fmt.Println("Driver is already on a ride")
		return
	}
	d.Available = false
	fmt.Printf("Driver %s started a ride\n", d.Name)
}

// EndRide marks the driver as available
func (d *Driver) EndRide() {
	if d.Available {
		fmt.Println("Driver is not on a ride")
		return
	}
	d.Available = true
	fmt.Printf("Driver %s ended the ride\n", d.Name)
}

// IsAvailable checks if the driver is available for a ride
func (d *Driver) IsAvailable() bool {
	return d.Available
}

// UpdateLocation updates the driver's location
func (d *Driver) UpdateLocation(latitude, longitude float64) {
	d.Location = NewLocation(latitude, longitude)
	fmt.Printf("Driver %s location updated to (%f, %f)\n", d.Name, latitude, longitude)
}
