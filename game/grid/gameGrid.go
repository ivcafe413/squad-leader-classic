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
	Label	string
	// spaces		[]spaces.GridSpace
	//hexMap	map[spaces.HexCoordinates]*spaces.Hex
	Map map[spaces.Hex]any
}

// type GridPopulation[V any] interface {
// 	populate (a spaces.GridSpace) V
// }

func New() *GameGrid {
	grid := new(GameGrid)
	grid.layers = []GridLayer {
		{
			Label: "default",
			// spaces: []spaces.GridSpace { },
			//hexMap: make(map[spaces.HexCoordinates]*spaces.Hex),
			//testMap: make(map[spaces.Hex]string),
			Map: GenerateRectangularMap(32, 9),
		},
	}

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