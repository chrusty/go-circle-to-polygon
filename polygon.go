package circletopolygon

import (
	"fmt"

	geojson "github.com/paulmach/go.geojson"
)

// ErrPolygonNotClosed is returned when the first and last points don't match:
var ErrPolygonNotClosed = fmt.Errorf("Polygons must start and finish on the same point")

// Polygon is a shape made up of points:
type Polygon []*Point

// GeoJSON renders a polygon as a GeoJSON geometry:
func (p Polygon) GeoJSON() ([]byte, error) {

	// Build up a slice of points for the GeoJSON library:
	geometryPoints := [][]float64{}
	for _, point := range p {
		geometryPoints = append(geometryPoints, []float64{
			float64(point.Longitude),
			float64(point.Latitude),
		})
	}

	// Use them to make a GeoJSON polygon:
	geoJSONPolygon := geojson.NewPolygonGeometry([][][]float64{
		geometryPoints,
	})

	return geoJSONPolygon.MarshalJSON()
}

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
