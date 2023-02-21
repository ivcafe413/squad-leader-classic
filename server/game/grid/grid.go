package grid

import (
	"encoding/json"
	//"fmt"

	"github.com/vagrant-technology/squad-leader/game/grid/coordinates"
	"github.com/vagrant-technology/squad-leader/game/grid/spaces"
)

type GameGrid[T coordinates.GridCoordinates[T]] interface {
	// 	NewSpace(system CoordinateSystem, layer int, coordinates ...int)
	Get(gc T) (*spaces.GridSpace, bool)
	Set(gc T, gs *spaces.GridSpace)
	SetMapGenerationStrat(MapGenerationStrategy[T])
}

// ----- Structs -----
type HexMap map[coordinates.HexCoordinates]*spaces.GridSpace

// TODO: Define Marshaling interface implementation
func (hm HexMap) MarshalJSON() ([]byte, error) {
	//return json.Marshal(hm)
	jsonMap := make(map[string]spaces.GridSpace)
	for k, v := range hm {
		jsonKey, _ := k.MarshalText()
		jsonMap[string(jsonKey)] = *v
	}

	return json.Marshal(jsonMap)
}

type HexGrid struct {
	//layers	[]GridLayer
	HexMap          `json:"map"`
	generationStrat MapGenerationStrategy[coordinates.HexCoordinates] `json:"-"`
}

// ----- HexGrid implementation of GameGrid interface -----
// Getter/Setter for embedded Map
func (hexGrid *HexGrid) Get(hc coordinates.HexCoordinates) (*spaces.GridSpace, bool) {
	result, exists := hexGrid.HexMap[hc]
	return result, exists
}

func (hexGrid *HexGrid) Set(hc coordinates.HexCoordinates, gs *spaces.GridSpace) {
	hexGrid.HexMap[hc] = gs
}

func (hexGrid *HexGrid) SetMapGenerationStrat(strat MapGenerationStrategy[coordinates.HexCoordinates]) {
	hexGrid.generationStrat = strat
}

// type GridLayer struct {
// 	Label	string
// }

func NewHexGrid(c, r int) *HexGrid {
	grid := new(HexGrid)

	grid.HexMap = make(map[coordinates.HexCoordinates]*spaces.GridSpace)
	grid.generationStrat = &RectangularHexMapStrat{
		RectangularMapStrat[coordinates.HexCoordinates]{
			columns: c,
			rows:    r,
		},
	}

	//Run Map Generation/Population Strategy
	grid.MakeMap()

	return grid
}

// Map generation (strategies?)
func (hexGrid *HexGrid) MakeMap() {
	hexGrid.generationStrat.GenerateMap(hexGrid)
}
