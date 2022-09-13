package grid

import (
	"testing"
	//"github.com/vagrant-technology/squad-leader/game/grid/spaces"
	"github.com/vagrant-technology/squad-leader/game/grid/coordinates"
)

func Test_GeneratedRectangleMap(t *testing.T) {
	var testGrid GameGrid[coordinates.HexCoordinates] //Testing interface

	testGrid = NewHexGrid()
	testGrid.GenerateRectangularMap(3, 3) //Should be 4x4 size

	//Top Left Corner - (0, 0, 0)
	topLeftHex, _ := coordinates.NewHex(0, 0)
	if _, exists := testGrid.Get(*topLeftHex); !exists {
		t.Fatalf(`Rectangle Map Generation failure - origin hex not existing`)
	}
	//Top Right Corner - (3, -1, -2)
	topRightHex, _ := coordinates.NewHex(3, -1)
	if _, exists := testGrid.Get(*topRightHex); !exists {
		t.Fatalf(`Rectangle Map Generation failure - top right hex not existing`)
	}
	//Bottom Left Corner - (0, 3, -3)
	bottomLeftHex, _ := coordinates.NewHex(0, 3)
	if _, exists := testGrid.Get(*bottomLeftHex); !exists {
		t.Fatalf(`Rectangle Map Generation failure - bottom left hex not existing`)
	}
	//Bottom Right Corner - (3, 2, -5)
	bottomRightHex, _ := coordinates.NewHex(3, 2)
	if _, exists := testGrid.Get(*bottomRightHex); !exists {
		t.Fatalf(`Rectangle Map Generation failure - bottom right hex not existing`)
	}

	nonExsistentHex, _ := coordinates.NewHex(100, 100)
	if _, exists := testGrid.Get(*nonExsistentHex); exists {
		t.Fatalf(`Rectangle Map Generation failure - (100, 100) should not exist`)
	}
}