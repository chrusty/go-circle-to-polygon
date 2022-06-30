package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBearing(t *testing.T) {
	assert.Equal(t, float64(0), Bearing(0, 12))
	assert.Equal(t, float64(-0.5235987755982988), Bearing(1, 12))
	assert.Equal(t, float64(-0.-0.7853981633974483), Bearing(4, 32))
}
