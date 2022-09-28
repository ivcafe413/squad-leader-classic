package coordinates

import (
	"fmt"
	"math"
	"errors"
)

// Abstract/Alias Direction to specific direction for interface compatibility
type Direction int

// ----- Interface -----
type GridCoordinates[T any] interface {
	//Equals(a, b *GridCoordinates) bool
	//Print(a GridCoordinates)
	Add(b T) T
	Subtract(b T) T
	Multiply(b int) T
	Distance(b T) int
	Neighbor(d Direction) T
}

// ----- Const/Enum -----
type HexDirection Direction
const (
	Undefined Direction = iota
	NorthEast
	SouthEast
	South
	SouthWest
	NorthWest
	North
)

var hexDirections = map[Direction]HexCoordinates {
	Undefined: { 0, 0, 0 },
	NorthEast: { 1, -1, 0 },
	SouthEast: { 1, 0, -1 },
	South: { 0, 1, -1 },
	SouthWest: { -1, 1, 0 },
	NorthWest: { -1, 0, 1 },
	North: { 0, -1, 1 },
}

// ----- Implementing Type -----
type HexCoordinates struct {
	q, r, s int
}

func (hc HexCoordinates) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%d, %d, %d", hc.q, hc.r, hc.s)), nil
}

// Need to use pointer on constructors
func NewHex(c ...int) (HexCoordinates, error) {
	//Only accepts coordinates or 2-3 dimensions
	if len(c) < 2 || len(c) > 3 {
		return HexCoordinates{-1, -1, -1}, errors.New("incorrect coordinate size for Hex")
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

	return r, nil
}

// func (a HexCoordinates) Equals(b HexCoordinates) bool {
// 	return a == b
// }

func (a HexCoordinates) Add(b HexCoordinates) HexCoordinates {
	//size := len(a.coordinates)
	var c HexCoordinates 
	//s := new(Hex)
	
	c.q = a.q + b.q
	c.r = a.r + b.r
	c.s = a.s + b.s

	//s.Coordinates = sCoords
	return c
}

func (a HexCoordinates) Subtract(b HexCoordinates) HexCoordinates {
	var c HexCoordinates 
	//s := new(Hex)

	c.q = a.q - b.q
	c.r = a.r - b.r
	c.s = a.s - b.s

	//s.Coordinates = sCoords
	return c
}

func (a HexCoordinates) Multiply(b int) HexCoordinates {
	var c HexCoordinates 
	//c := new(Hex)

	c.q = a.q * b
	c.r = a.r * b
	c.s = a.s * b

	return c
}

func (a HexCoordinates) Distance(b HexCoordinates) int {
	c := a.Subtract(b)

	//The following only works on CUBIC coordinates, which this is under the hood
	d := 0.0

	d += math.Abs(float64(c.q))
	d += math.Abs(float64(c.r))
	d += math.Abs(float64(c.s))

	d /= 2

	return int(d)
}

func (a HexCoordinates) Neighbor(d Direction) HexCoordinates {
	return a.Add(hexDirections[d])
}