package grid

import (
	"github.com/vagrant-technology/squad-leader/game/grid/spaces"
)

//Define constants/enums
type CoordinateSystem int
const (
	Undefined CoordinateSystem = iota
	Axial
	Cubic
	Offset
)

// type gameGrid interface {
// 	NewSpace(system CoordinateSystem, layer int, coordinates ...int)
// }

// ----- Structs -----

type GameGrid struct {
	layers	[]GridLayer
}

type GridLayer struct {
	label		string
	spaces		[]spaces.GridSpace
}

func New() *GameGrid {
	grid := new(GameGrid)
	grid.layers = []GridLayer {
		{
			label: "default",
			spaces: []spaces.GridSpace { },
		},
	}

	return grid
}