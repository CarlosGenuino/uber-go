package domain

import "math"

type Location struct {
	Latitude  float64
	Longitude float64
}

func NewLocation(latitude, longitude float64) *Location {
	return &Location{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (l *Location) DistanceTo(other *Location) float64 {
	const earthRadius = 6371 // Earth's radius in kilometers

	lat1 := l.Latitude * math.Pi / 180
	lon1 := l.Longitude * math.Pi / 180
	lat2 := other.Latitude * math.Pi / 180
	lon2 := other.Longitude * math.Pi / 180

	dLat := lat2 - lat1
	dLon := lon2 - lon1
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
