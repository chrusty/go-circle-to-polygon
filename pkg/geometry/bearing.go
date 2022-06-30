package geometry

import (
	"math"
)

func Bearing(i, n int) float64 {
	const start float64 = 0
	const direction float64 = 1

	return start + direction*float64(2)*math.Pi*float64(-i)/float64(n)
}
