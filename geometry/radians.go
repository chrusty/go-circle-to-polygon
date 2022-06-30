package geometry

import "math"

func Radians(angleInDegrees float64) float64 {
	return (angleInDegrees * math.Pi) / 180
}
