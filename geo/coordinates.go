// Package geo for coordinate defination
package geo

// Coordinates struct type
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// SetLatitude sets the latitude for Coordinate struct
func (c *Coordinates) SetLatitude(latitude float64) {
	c.Latitude = latitude
}

// SetLongitude sets the longitude for Coordinate struct
func (c *Coordinates) SetLongitude(longitude float64) {
	c.Longitude = longitude
}
