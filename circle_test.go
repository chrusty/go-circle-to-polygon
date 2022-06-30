package circletopolygon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle(t *testing.T) {

	// Make a circle:
	circle := &Circle{
		Centre: &Point{
			Latitude:  -41.270634,
			Longitude: 173.283966,
		},
		Radius: 20000,
	}

	// Convert it to a Polygon with 60 edges:
	polygon, err := circle.ToPolygon(60)
	assert.NoError(t, err)
	assert.Len(t, polygon, 61)

	// Render as GeoJSON:
	geoJSON, err := polygon.GeoJSON()
	assert.NoError(t, err)
	assert.Len(t, geoJSON, 2453)
}
