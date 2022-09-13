package spaces

import (
	//"errors"
	"github.com/vagrant-technology/squad-leader/game/grid/pieces"
	"github.com/vagrant-technology/squad-leader/game/grid/spaces/coordinates"
)

type GridSpace struct {
	pieces []pieces.GridPiece
}

type Hex struct {
	GridSpace
	// Private/un-exported fields

	// Instead now we are EMBEDDING the interface
	coordinates.GridCoordinates[coordinates.HexCoordinates]
}

func NewHex(c ...int) (*Hex, error) {
	var pieces []pieces.GridPiece

	hexCoordinates, err := coordinates.NewHex(c...)
	newHex := Hex {
		GridSpace{ pieces },
		hexCoordinates,
	}

	return &newHex, err
}

// ----- Hex implementation of GridSpace interface -----
