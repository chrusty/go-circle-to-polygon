package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	assert.Equal(t, float64(0.0003135711885774796), DistanceByRadius(2000))
	assert.Equal(t, float64(3.135711885774796), DistanceByRadius(20000000))
}
