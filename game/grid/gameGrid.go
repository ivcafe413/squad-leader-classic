package grid

import (
	"github.com/vagrant-technology/squad-leader/game/grid/spaces"
	"github.com/vagrant-technology/squad-leader/game/grid/coordinates"
)

type GameGrid[T coordinates.GridCoordinates[T]] interface {
// 	NewSpace(system CoordinateSystem, layer int, coordinates ...int)
	Get(gc T) (*spaces.GridSpace, bool)
	Set(gc T, gs *spaces.GridSpace)
	GenerateRectangularMap(right, bottom int)
}

// ----- Structs -----
type hexMap map[coordinates.HexCoordinates]*spaces.GridSpace

type HexGrid struct {
	//layers	[]GridLayer
	hexMap
}

// Getter/Setter for embedded Map
func (hexGrid HexGrid) Get(hc coordinates.HexCoordinates) (*spaces.GridSpace, bool) {
	result, exists := hexGrid.hexMap[hc]
	return result, exists
}

func (hexGrid HexGrid) Set(hc coordinates.HexCoordinates, gs *spaces.GridSpace) {
	hexGrid.hexMap[hc] = gs
}

type GridLayer struct {
	Label	string	
}

// type GridPopulation[V any] interface {
// 	populate (a spaces.GridSpace) V
// }

func NewHexGrid() *HexGrid {
	grid := new(HexGrid)
	grid.hexMap = make(map[coordinates.HexCoordinates]*spaces.GridSpace)
	return grid
}

//Map generation (strategies?)
func (hexGrid *HexGrid) GenerateRectangularMap (right, bottom int) {
	//right and bottom dictate furthest coordinate, not width/height

	// newMap := make(map[*spaces.GridSpace]V)
	// var newHex *spaces.GridSpace

	//Normal map-right is 32 (33 spaces wide)
	for i := 0; i <= right; i++ {
		q_offset := i >> 1
		//Normal map-bottom is 9 (10 spaces wide)
		for j := -q_offset; j <= bottom - q_offset; j++ {
			//newHex, _ = spaces.NewHex(i, j)
			newHex, _ := coordinates.NewHex(i, j)
			newSpace := spaces.New()
			// newMap[*newHex] = make(V)
			//hexGrid.Map[newHex] = newSpace
			//hexGrid.HexMap[newHex] = newSpace
			hexGrid.Set(*newHex, newSpace)
		}
	}

	//return newMap
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