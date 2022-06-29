package circletopolygon

// Circle represents a circular area:
type Circle struct {
	Centre *Point `json:"centre"`
	Radius int32  `json:"radius"`
}

// ToPolygon renders a polygon (with the given number of edges) to represent a circle:
func (c *Circle) ToPolygon(edges int32) (*Polygon, error) {

	return nil, nil
}
