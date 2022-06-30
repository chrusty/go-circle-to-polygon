package main

import (
	"flag"
	"fmt"

	circletopolygon "github.com/chrusty/go-circle-to-polygon"
)

var (
	edges     = flag.Int64("edges", 33, "number of edges")
	latitude  = flag.Float64("lat", -41.270634, "latitude")
	longitude = flag.Float64("lon", 173.283966, "longitude")
	radius    = flag.Int64("radius", 20000, "radius in metres")
)

func main() {

	flag.Parse()

	// Make a circle:
	circle := &circletopolygon.Circle{
		Centre: &circletopolygon.Point{
			Latitude:  float32(*latitude),
			Longitude: float32(*longitude),
		},
		Radius: int32(*radius),
	}

	// Validate the circle:
	if err := circle.Validate(); err != nil {
		panic(err)
	}

	// Convert it to a Polygon with 10 edges:
	polygon, err := circle.ToPolygon(int(*edges))
	if err != nil {
		panic(err)
	}

	// Render as GeoJSON:
	geoJSON, err := polygon.GeoJSON()
	if err != nil {
		panic(err)
	}

	// Print the result:
	fmt.Println(string(geoJSON))
}
