package spaces

import (
	"github.com/vagrant-technology/squad-leader/game/grid/pieces"
)

type GridSpace struct {
	Name string `json:"name"`
	Pieces []pieces.GridPiece `json:"pieces"`
}

func New(n string) *GridSpace {
	//var pieces []pieces.GridPiece
	
	newSpace := new(GridSpace)
	newSpace.Pieces = make([]pieces.GridPiece, 0)//pieces

	newSpace.Name = n

	return newSpace
}
