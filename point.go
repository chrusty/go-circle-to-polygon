package circletopolygon

import "fmt"

// ErrPointOutOfBounds is returned a latitude or longitude is outside of the -180 to +180 range:
var ErrPointOutOfBounds = fmt.Errorf("Latitudes and longitudes must be between -180 and +180")

// Point represents a geographic coordinate:
type Point struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// Validate that our Point makes sense:
func (p *Point) Validate() error {

	// Check that the latitudes and longitudes are within sensible limits:
	if p.Latitude < -180 || p.Latitude > 180 || p.Longitude < -180 || p.Longitude > 180 {
		return ErrPointOutOfBounds
	}

	return nil
}
