package spaces

import (
	"math"
	"errors"
)

type HexDirection int
const (
	Undefined HexDirection = iota
	NorthEast
	SouthEast
	South
	SouthWest
	NorthWest
	North
)

var directionMap = map[HexDirection]*Hex {
	Undefined: { coordinates: [3]int {0, 0, 0} },
	NorthEast: { coordinates: [3]int {1, -1, 0} },
	SouthEast: { coordinates: [3]int {1, 0, -1} },
	South: { coordinates: [3]int {0, 1, -1} },
	SouthWest: { coordinates: [3]int {-1, 1, 0} },
	NorthWest: { coordinates: [3]int {-1, 0, 1} },
	North: { coordinates: [3]int {0, -1, 1} },
}

func NewHex(c ...int) (*Hex, error) {
	//Only accepts coordinates or 2-3 dimensions
	if len(c) < 2 || len(c) > 3 {
		return nil, errors.New("incorrect coordinate size for Hex")
	}

	//Store 3rd axial coordinate if not already present
	var s int
	if len(c) == 2 {
		// c = c[0:3]
		s = -c[0] - c[1]
	} else {
		s = c[3]
	}

	r := [3]int {c[0], c[1], s}

	hex := &Hex {
		coordinates: r,
	}
	return hex, nil
}

// ----- Interfaces -----
type GridSpace interface {
	Equals(a, b *GridSpace) bool
	Add(a, b *GridSpace) *GridSpace
	Subtract(a, b *GridSpace) *GridSpace
	Multiply(a *GridSpace, b int) *GridSpace
	Distance(a, b *GridSpace) int
	Neighbor(a *GridSpace, d HexDirection) *GridSpace
}

// ----- Structs -----
type Hex struct {
	coordinates	[3]int
}

// ----- Interface Implementations -----
// ----- Hex -> GridSpace -----
func (a *Hex) Equals(b *Hex) bool {
	// Short-circuit on falses to test equality
	if len(a.coordinates) != len(b.coordinates) {
		return false
	}

	for i := range a.coordinates {
		if a.coordinates[i] != b.coordinates[i] {
			return false
		}
	}

	return true
}

func (a *Hex) Add(b *Hex) *Hex {
	//size := len(a.coordinates)
	var sCoords [3]int 
	s := new(Hex)
	s.coordinates = sCoords

	for i := range a.coordinates {
		s.coordinates[i] = a.coordinates[i] + b.coordinates[i]
	}
	return s
}

func (a *Hex) Subtract(b *Hex) *Hex {
	var sCoords [3]int 
	s := new(Hex)
	s.coordinates = sCoords

	for i := range a.coordinates {
		s.coordinates[i] = a.coordinates[i] - b.coordinates[i]
	}

	return s
}

func (a *Hex) Multiply(b int) *Hex {
	var cCoords [3]int 
	c := new(Hex)
	c.coordinates = cCoords

	for i := range a.coordinates {
		c.coordinates[i] = a.coordinates[i] * b
	}

	return c
}

func (a *Hex) Distance(b *Hex) int {
	c := a.Subtract(b)

	//The following only works on CUBIC coordinates, which this is under the hood
	d := 0.0
	for _,n := range(c.coordinates) {
		d += math.Abs(float64(n))
	}
	d /= 2

	return int(d)
}

func (a *Hex) Neighbor(d HexDirection) *Hex {
	return a.Add(directionMap[d])
}