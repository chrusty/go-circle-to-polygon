package circletopolygon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {

	// Make a point with an invalid latitude:
	point := &Point{
		Latitude:  -181,
		Longitude: 0,
	}
	assert.Error(t, point.Validate())

	// Fix the latitude, break the longitude:
	point.Latitude = -41.270634
	point.Longitude = 181
	assert.Error(t, point.Validate())

	// Fix the longitude:
	point.Longitude = 173.283966
	assert.NoError(t, point.Validate())
}
