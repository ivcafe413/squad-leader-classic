package grid

import (
	"github.com/vagrant-technology/squad-leader/game/grid/spaces"
)

//Define constants/enums
// type CoordinateSystem int
// const (
// 	Undefined CoordinateSystem = iota
// 	Axial
// 	Cubic
// 	Offset
// )

// type gameGrid interface {
// 	NewSpace(system CoordinateSystem, layer int, coordinates ...int)
// }

// ----- Structs -----

type GameGrid struct {
	layers	[]GridLayer
}

type GridLayer struct {
	label	string
	// spaces		[]spaces.GridSpace
	//hexMap	map[spaces.HexCoordinates]*spaces.Hex
	testMap map[spaces.Hex]string
}

func New() *GameGrid {
	grid := new(GameGrid)
	grid.layers = []GridLayer {
		{
			label: "default",
			// spaces: []spaces.GridSpace { },
			//hexMap: make(map[spaces.HexCoordinates]*spaces.Hex),
			//testMap: make(map[spaces.Hex]string),
			testMap: GenerateRectangularMap(32, 9),
		},
	}

	return grid
}

//Map generation (strategies?)
func GenerateRectangularMap(right, bottom int) map[spaces.Hex]string{
	//right and bottom dictate furthest coordinate, not width/height

	newMap := make(map[spaces.Hex]string)
	var newHex *spaces.Hex

	//Normal map-right is 32 (33 spaces wide)
	for i := 0; i <= right; i++ {
		q_offset := i >> 1
		//Normal map-bottom is 9 (10 spaces wide)
		for j := -q_offset; j <= bottom - q_offset; j++ {
			newHex, _ = spaces.NewHex(i, j)
			newMap[*newHex] = "Test Generate"
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