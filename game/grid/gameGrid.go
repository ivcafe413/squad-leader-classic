package grid

import (
	"github.com/vagrant-technology/squad-leader/game/grid/spaces"
	"github.com/vagrant-technology/squad-leader/game/grid/spaces/coordinates"
)

// type gameGrid interface {
// 	NewSpace(system CoordinateSystem, layer int, coordinates ...int)
// }

// ----- Structs -----

type HexGrid struct {
	//layers	[]GridLayer
	Map map[coordinates.GridCoordinates[coordinates.HexCoordinates]]spaces.Hex
}

type GridLayer struct {
	Label	string
	// spaces		[]spaces.GridSpace
	//hexMap	map[spaces.HexCoordinates]*spaces.Hex
	
}

// type GridPopulation[V any] interface {
// 	populate (a spaces.GridSpace) V
// }

func NewHexGrid() *HexGrid {
	grid := new(HexGrid)
	// grid.layers = []GridLayer {
	// 	{
	// 		Label: "default",
	// 		Map: GenerateRectangularMap[string](32, 9),
	// 	},
	// }
	

	return grid
}

//Map generation (strategies?)
func GenerateRectangularMap[V any] (right, bottom int) map[spaces.Hex]V {
	//right and bottom dictate furthest coordinate, not width/height

	newMap := make(map[spaces.Hex]V)
	var newHex *spaces.Hex

	//Normal map-right is 32 (33 spaces wide)
	for i := 0; i <= right; i++ {
		q_offset := i >> 1
		//Normal map-bottom is 9 (10 spaces wide)
		for j := -q_offset; j <= bottom - q_offset; j++ {
			newHex, _ = spaces.NewHex(i, j)
			newMap[*newHex] = make(V)
		}
	}

	return newMap
}

func PopulateRectangularMap(gameMap map[spaces.Hex]string, right, bottom int) {
	//right and bottom dictate furthest coordinate, not width/height
	var newHex *spaces.Hex
	for i := 0; i <= right; i++ { //Normal map-right is 32 (33 spaces wide)
		q_offset := i >> 1
		for j := -q_offset; j <= bottom - q_offset; j++ { //Normal map-bottom is 9 (10 spaces wide)
			newHex, _ = spaces.NewHex(i, j)
			gameMap[*newHex] = "Test Populate"
		}
	}
}

// func 