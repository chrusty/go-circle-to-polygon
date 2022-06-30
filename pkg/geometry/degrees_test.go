package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDegrees(t *testing.T) {
	assert.Equal(t, float64(-41.27063400000001), Degrees(-0.7203084476855175))
	assert.Equal(t, float64(173.283966), Degrees(3.024375747613908))
}
