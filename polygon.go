package circletopolygon

import "fmt"

// ErrPolygonNotClosed is returned when the first and last points don't match:
var ErrPolygonNotClosed = fmt.Errorf("Polygons must start and finish on the same point")

// Polygon is a shape made up of points:
type Polygon []*Point

// Validate that our Polygon makes sense:
func (p Polygon) Validate() error {

	// Get the points:
	points := p

	// Make sure the first point is the same as the last (that we have a "closed" geometry):
	if points[0].Latitude != points[len(points)-1].Latitude || points[0].Longitude != points[len(points)-1].Longitude {
		return ErrPolygonNotClosed
	}

	// Validate all of the points:
	for _, point := range points {
		if err := point.Validate(); err != nil {
			return err
		}
	}

	return nil
}
