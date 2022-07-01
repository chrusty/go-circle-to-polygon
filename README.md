Circle-to-Polygon
=================

A Go library to convert a circle to a polygon (port of https://github.com/gabzim/circle-to-polygon).


![Example](example.png "Example circle (20km radius with 33 edges)")

Usage
-----

The library offers some basic geometry components:

- [Point](point.go): A geographic coordinate (lat/lon)
- [Circle](circle.go): A point with a radius (in metres)
- [Polygon](polygon.go): A list of points

Each of these structures offers a `Validate()` method which returns an error if the data doesn't make sense (lat/lon out of bounds, or an open polygon).


A working example can be found under [cmd/circle](cmd/circle/main.go). This renders a custom circle as GeoJSON, which you can then load into https://geojson.io to view!


### Make a circle

```go
    circle := &circletopolygon.Circle{
        Centre: &circletopolygon.Point{
            Latitude:  -41.270634,
            Longitude: 173.283966,
        },
        Radius: 20000, // 20km
    }
```


### Convert a circle to a polygon with 32 edges:

```go
    polygon, err := circle.ToPolygon(32)
    if err != nil {
        panic(err)
    }
```


### Render a polygon as GeoJSON:

```go
    geoJSON, err := polygon.GeoJSON()
    if err != nil {
        panic(err)
    }
```


### Do other things with the polygon:

```go
    for _, point := range polygon {
        fmt.Printf("Lat, Lon: %f, %f", point.Latitude, point.Longitude)
    }
```
