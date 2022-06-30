package geometry

func DistanceByRadius(distanceInMetres float64) float64 {
	const earthRadiusMetres = 6378137

	return distanceInMetres / earthRadiusMetres
}
