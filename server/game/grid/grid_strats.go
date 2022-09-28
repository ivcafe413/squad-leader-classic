package grid

import (
	"github.com/vagrant-technology/squad-leader/game/grid/coordinates"
	"github.com/vagrant-technology/squad-leader/game/grid/spaces"
)

type MapGenerationStrategy[T coordinates.GridCoordinates[T]] interface {
	GenerateMap(GameGrid[T])
}

type RectangularMapStrat[T coordinates.GridCoordinates[T]] struct {
	columns, rows int
}

type RectangularHexMapStrat struct {
	RectangularMapStrat[coordinates.HexCoordinates] //This is actually very nice
}

func (s *RectangularHexMapStrat) GenerateMap(gg GameGrid[coordinates.HexCoordinates]) {
	//MakeMap should overwrite any previous map completely (no orphaned coordinates/spaces/pieces)

	//right and bottom dictate furthest q/r coordinates, not width/height

	//Normal map-right is 32 (33 columns wide)
	for i := 0; i < s.columns; i++ {
		q_offset := i >> 1
		//Normal map-bottom is 9 (10 rows tall)
		for j := -q_offset; j < s.rows - q_offset; j++ {
			newHex, _ := coordinates.NewHex(i, j)
			//newSpace := spaces.New()
			
			gg.Set(newHex, spaces.New("placeholder"))
		}
	}
}