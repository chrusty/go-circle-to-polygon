package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadians(t *testing.T) {
	assert.Equal(t, float64(-0.7203084476855175), Radians(-41.270634))
	assert.Equal(t, float64(3.024375747613908), Radians(173.283966))
}
