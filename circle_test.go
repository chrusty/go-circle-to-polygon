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
		Radius: 2000,
	}

	// Convert it to a Polygon with 6 edges:
	polygon, err := circle.ToPolygon(6)
	assert.NoError(t, err)
	assert.Len(t, polygon, 7)
}
