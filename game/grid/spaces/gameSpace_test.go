package spaces

import (
	"testing"
)

func Test_HexEquals(t *testing.T) {
	// hexA := &Hex {
	// 	coordinates: []int {1, 1},
	// }
	// hexB := &Hex {
	// 	coordinates: []int {1, 1},
	// }
	//coords := []int {1, 1}
	hexA, _ := NewHex(1, 1)
	hexB, _ := NewHex(1, 1)

	if !hexA.Equals(hexB) {
		t.Fatalf(`Hex .Equals Test failing`)
	}

	// -----REMOVED, BAD TEST: == doesn't ever work with different pointers
	// if!(hexA == hexB) {
	// 	t.Fatalf(`Hex == (equality) Test failing: %v should equal %v`, hexA, hexB)
	// }

	//coords = []int {2, 2} //NOT equal
	hexB, _ = NewHex(2, 2)

	if hexA.Equals(hexB) {
		t.Fatalf(`Hex Equality Test failing - (1, 1) should NOT equal (2, 2)`)
	}

	// if hexA == hexB {
	// 	t.Fatalf(`Hex == (equality operator) Test failing `)
	// }
}

func Test_HexAdd(t *testing.T) {
	// hexA := &Hex {
	// 	coordinates: []int {2,0},
	// }
	// hexB := &Hex {
	// 	coordinates: []int {0,-1},
	// }
	//coordsA := []int {2, 0}
	//coordsB := []int {0, -1}
	hexA, _ := NewHex(2, 0)
	hexB, _ := NewHex(0, -1)

	hexC := hexA.Add(hexB)
	// hexD := &Hex {
	// 	coordinates: []int {2, -1},
	// }
	hexD, _ := NewHex(2, -1)
	if !hexC.Equals(hexD) {
		t.Fatalf(`Hex Addition Failing - (2, 0) + (0, -1) SHOULD = (2, -1)`)
	}
}

func Test_HexSubtract(t *testing.T) {
	hexA, _ := NewHex(2, 1)
	hexB, _ := NewHex(2, -1)

	hexC := hexA.Subtract(hexB)
	hexR, _ := NewHex(0, 2)

	if !hexC.Equals(hexR) {
		t.Fatalf(`Hex Subtraction Failing - (2, 1) - (2, -1) SHOULD = (0, 2)`)
	}
}

func Test_HexMultiply(t *testing.T) {
	hexA, _ := NewHex(2, 0)

	hexM := hexA.Multiply(2)
	hexR, _ := NewHex(4, 0)

	if !hexM.Equals(hexR) {
		t.Fatalf(`Hex Multiplication Failing`)
	}
}

func Test_HexDistance(t *testing.T) {
	hexA, _ := NewHex(2, 2) //s = -4, distance is 4 from 0
	hexB, _ := NewHex(0, 0)
	hexC, _ := NewHex(6, 0) //distance is 4 from hexA, 6 from 0, s = -6
	hexD, _ := NewHex(3, -1) //s = -2, distance is 3 from 0

	testDistance := hexA.Distance(hexB)
	if testDistance != 4 {
		t.Fatalf(`Hex Distance algorithm failing - distance from (2, 2) -> (0, 0) should be 4, not %v`, testDistance)
	}

	testDistance = hexA.Distance(hexC)
	if testDistance != 4 {
		t.Fatalf(`Hex Distance algorithm failing - distance from (2, 2) -> (6, 0) should be 4, not %v`, testDistance)
	}

	testDistance = hexB.Distance(hexD)
	if testDistance != 3 {
		t.Fatalf(`Hex Distance algorithm failure #3 - distance from (0, 0) -> (3, -1) should be 3, not %v`, testDistance)
	}
}

func Test_HexDirection(t *testing.T) {
	hexA, _ := NewHex(0, 0)
	hexSouth := hexA.Neighbor(South) //Should be (0, 1)

	testHex, _ := NewHex(0, 1)
	if !hexSouth.Equals(testHex) {
		t.Fatalf(`Hex Neighbor failure - new hex should be (0, 1)`)
	}
}