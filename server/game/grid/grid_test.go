package grid

import (
	"testing"
	"fmt"
	//"github.com/vagrant-technology/squad-leader/game/grid/spaces"
	"github.com/vagrant-technology/squad-leader/game/grid/coordinates"
)

func Test_GeneratedRectangleMap(t *testing.T) {
	var testGrid GameGrid[coordinates.HexCoordinates]
	//Ignore staticcheck, declaration and assignment separated to test interface
	testGrid = NewHexGrid(33, 10)
	//fmt.Println(testGrid.hexMap)
	// for key, value := range *testGrid { // Order not specified 
	// 	fmt.Println(key, value)
	// }

	//Top Left Corner - (0, 0, 0)
	topLeftHex, _ := coordinates.NewHex(0, 0)
	topLeftGrid, exists := testGrid.Get(topLeftHex)
	fmt.Printf("(0, 0): %v\n", topLeftGrid)
	if !exists {
		t.Fatalf(`Rectangle Map Generation failure - origin hex not existing`)
	}
	//Top Right Corner - (3, -1, -2)
	topRightHex, _ := coordinates.NewHex(3, -1)
	if _, exists := testGrid.Get(topRightHex); !exists {
		t.Fatalf(`Rectangle Map Generation failure - top right hex not existing`)
	}
	//Bottom Left Corner - (0, 3, -3)
	bottomLeftHex, _ := coordinates.NewHex(0, 3)
	if _, exists := testGrid.Get(bottomLeftHex); !exists {
		t.Fatalf(`Rectangle Map Generation failure - bottom left hex not existing`)
	}
	//Bottom Right Corner - (3, 2, -5)
	bottomRightHex, _ := coordinates.NewHex(3, 2)
	if _, exists := testGrid.Get(bottomRightHex); !exists {
		t.Fatalf(`Rectangle Map Generation failure - bottom right hex not existing`)
	}

	nonExsistentHex, _ := coordinates.NewHex(100, 100)
	if _, exists := testGrid.Get(nonExsistentHex); exists {
		t.Fatalf(`Rectangle Map Generation failure - (100, 100) should not exist`)
	}
}