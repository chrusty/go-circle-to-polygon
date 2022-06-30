package geometry

import "math"

func Degrees(angleInRadians float64) float64 {
	return (angleInRadians * 180) / math.Pi
}
