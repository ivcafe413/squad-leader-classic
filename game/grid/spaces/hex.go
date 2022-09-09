package spaces

import (
	"math"
	"errors"
)

type Hex struct {
	//Private/un-exported
	//coordinates	[3]int //TODO: Replace [3]int with HexCoordinates?
	Coordinates HexCoordinates
}

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

var hexDirections = map[HexDirection]*Hex {
	Undefined: { Coordinates: HexCoordinates {0, 0, 0} },
	NorthEast: { Coordinates: HexCoordinates {1, -1, 0} },
	SouthEast: { Coordinates: HexCoordinates {1, 0, -1} },
	South: { Coordinates: HexCoordinates {0, 1, -1} },
	SouthWest: { Coordinates: HexCoordinates {-1, 1, 0} },
	NorthWest: { Coordinates: HexCoordinates {-1, 0, 1} },
	North: { Coordinates: HexCoordinates {0, -1, 1} },
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

	r := HexCoordinates {c[0], c[1], s}

	hex := &Hex {
		Coordinates: r,
	}
	return hex, nil
}

// ----- Hex implementation of GridSpace interface -----
func (a *Hex) Equals(b *Hex) bool {
	return a.Coordinates == b.Coordinates
}

func (a *Hex) Add(b *Hex) *Hex {
	//size := len(a.coordinates)
	var sCoords HexCoordinates 
	s := new(Hex)
	
	sCoords.q = a.Coordinates.q + b.Coordinates.q
	sCoords.r = a.Coordinates.r + b.Coordinates.r
	sCoords.s = a.Coordinates.s + b.Coordinates.s

	s.Coordinates = sCoords
	return s
}

func (a *Hex) Subtract(b *Hex) *Hex {
	var sCoords HexCoordinates 
	s := new(Hex)

	sCoords.q = a.Coordinates.q - b.Coordinates.q
	sCoords.r = a.Coordinates.r - b.Coordinates.r
	sCoords.s = a.Coordinates.s - b.Coordinates.s

	s.Coordinates = sCoords
	return s
}

func (a *Hex) Multiply(b int) *Hex {
	var cCoords HexCoordinates 
	c := new(Hex)

	cCoords.q = a.Coordinates.q * b
	cCoords.r = a.Coordinates.r * b
	cCoords.s = a.Coordinates.s * b

	c.Coordinates = cCoords
	return c
}

func (a *Hex) Distance(b *Hex) int {
	c := a.Subtract(b)

	//The following only works on CUBIC coordinates, which this is under the hood
	d := 0.0

	d += math.Abs(float64(c.Coordinates.q))
	d += math.Abs(float64(c.Coordinates.r))
	d += math.Abs(float64(c.Coordinates.s))

	d /= 2

	return int(d)
}

func (a *Hex) Neighbor(d HexDirection) *Hex {
	return a.Add(hexDirections[d])
}