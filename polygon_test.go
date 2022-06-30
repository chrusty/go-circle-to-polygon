package circletopolygon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolygon(t *testing.T) {

	// Make an invalid polygon (start and end don't match):
	polygon := Polygon{
		{
			Latitude:  1,
			Longitude: 2,
		},
		{
			Latitude:  3,
			Longitude: 4,
		},
	}
	assert.Error(t, polygon.Validate())
	assert.Equal(t, ErrPolygonNotClosed, polygon.Validate())

	// Make the last point match the first:
	polygon = append(polygon, &Point{Latitude: 1, Longitude: 2})
	assert.NoError(t, polygon.Validate())

	// Include an invalid point:
	polygon[1] = &Point{Latitude: 181, Longitude: 2}
	assert.Error(t, polygon.Validate())
	assert.Equal(t, ErrPointOutOfBounds, polygon.Validate())
}
