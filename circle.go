package circletopolygon

// Circle represents a circular area:
type Circle struct {
	Centre *Point `json:"centre"` // The centre coordinates
	Radius int32  `json:"radius"` // The radius (in metres)
}

// ToPolygon renders a polygon (with the given number of edges) to represent a circle:
func (c *Circle) ToPolygon(edges int32) (Polygon, error) {

	// First validate the centre point:
	if err := c.Centre.Validate(); err != nil {
		return nil, err
	}

	var polygon Polygon
	return polygon, nil
}
