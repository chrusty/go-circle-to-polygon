package circletopolygon

import (
	"math"

	"github.com/chrusty/go-circle-to-polygon/pkg/geometry"
)

// Circle represents a circular area:
type Circle struct {
	Centre *Point `json:"centre"` // The centre coordinates
	Radius int32  `json:"radius"` // The radius (in metres)
}

// ToPolygon renders a polygon (with the given number of edges) to represent a circle:
func (c *Circle) ToPolygon(edges int) (Polygon, error) {

	// First validate the centre point:
	if err := c.Centre.Validate(); err != nil {
		return nil, err
	}

	// Make our polygon (with our centre as the first point):
	polygon := Polygon{}
	distanceByEarth := geometry.DistanceByRadius(float64(c.Radius))
	latAngle := geometry.Radians(float64(c.Centre.Latitude))
	lonAngle := geometry.Radians(float64(c.Centre.Longitude))

	// Iterate through the edges and calculate each point:
	for i := 0; i < edges; i++ {

		// Get the bearing to this point:
		bearing := geometry.Bearing(i, edges)

		// Calculate the latitude of the next point:
		latitude := math.Asin(math.Sin(latAngle)*math.Cos(distanceByEarth) + math.Cos(latAngle)*math.Sin(distanceByEarth)*math.Cos(bearing))

		// Calculate the longitude of the next point:
		longitude := lonAngle + math.Atan2(
			math.Sin(bearing)*math.Sin(distanceByEarth)*math.Cos(latAngle),
			math.Cos(distanceByEarth)-math.Sin(latAngle)*math.Sin(latitude),
		)

		// Add the next point to our polygon:
		polygon = append(polygon, &Point{
			Latitude:  float32(geometry.Degrees(latitude)),
			Longitude: float32(geometry.Degrees(longitude)),
		})
	}

	// Make sure the last point matches up with the first:
	polygon = append(polygon, polygon[0])

	// Make sure we ended up with a valid polygon:
	if err := polygon.Validate(); err != nil {
		return nil, err
	}

	return polygon, nil
}

// Validate that our Circle makes sense:
func (c *Circle) Validate() error {
	return c.Centre.Validate()
}
