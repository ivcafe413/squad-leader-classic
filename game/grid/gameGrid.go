package grid

import (
	"github.com/vagrant-technology/squad-leader/game/grid/spaces"
	"github.com/vagrant-technology/squad-leader/game/grid/coordinates"
)

type GameGrid[T coordinates.GridCoordinates[T]] interface {
// 	NewSpace(system CoordinateSystem, layer int, coordinates ...int)
	Get(gc T) (*spaces.GridSpace, bool)
	Set(gc T, gs *spaces.GridSpace)
	SetMapGenerationStrat(MapGenerationStrategy[T])
}

// ----- Structs -----
type hexMap map[coordinates.HexCoordinates]*spaces.GridSpace

type HexGrid struct {
	//layers	[]GridLayer
	hexMap
	generationStrat MapGenerationStrategy[coordinates.HexCoordinates]
}

// Getter/Setter for embedded Map
func (hexGrid *HexGrid) Get(hc coordinates.HexCoordinates) (*spaces.GridSpace, bool) {
	result, exists := hexGrid.hexMap[hc]
	return result, exists
}

func (hexGrid *HexGrid) Set(hc coordinates.HexCoordinates, gs *spaces.GridSpace) {
	hexGrid.hexMap[hc] = gs
}

func (hexGrid *HexGrid) SetMapGenerationStrat(strat MapGenerationStrategy[coordinates.HexCoordinates]) {
	hexGrid.generationStrat = strat
}

// type GridLayer struct {
// 	Label	string	
// }

func NewHexGrid(c, r int) *HexGrid {
	grid := new(HexGrid)

	grid.hexMap = make(map[coordinates.HexCoordinates]*spaces.GridSpace)
	grid.generationStrat = &RectangularHexMapStrat {
		RectangularMapStrat[coordinates.HexCoordinates] {
			columns: c,
			rows: r,
		},
	}
	
	//Run Map Generation/Population Strategy
	grid.MakeMap()
	
	return grid
}

//Map generation (strategies?)
func (hexGrid *HexGrid) MakeMap () {
	hexGrid.generationStrat.GenerateMap(hexGrid)
}

// func PopulateRectangularMap(gameMap map[spaces.Hex]string, right, bottom int) {
// 	//right and bottom dictate furthest coordinate, not width/height
// 	var newHex *spaces.Hex
// 	for i := 0; i <= right; i++ { //Normal map-right is 32 (33 spaces wide)
// 		q_offset := i >> 1
// 		for j := -q_offset; j <= bottom - q_offset; j++ { //Normal map-bottom is 9 (10 spaces wide)
// 			newHex, _ = spaces.NewHex(i, j)
// 			gameMap[*newHex] = "Test Populate"
// 		}
// 	}
// }

// func 