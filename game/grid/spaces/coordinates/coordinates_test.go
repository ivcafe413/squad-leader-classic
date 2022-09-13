package coordinates

import (
	"testing"
)

func Test_HexEquals(t *testing.T) {
	hexA, _ := NewHex(1, 1)
	hexB, _ := NewHex(1, 1)

	if *hexA != *hexB {
		t.Fatalf(`Hex equality Test failing`)
	}

	//coords = []int {2, 2} //NOT equal
	hexB, _ = NewHex(2, 2)

	if *hexA == *hexB {
		t.Fatalf(`Hex equality Test failing - (1, 1) should NOT equal (2, 2)`)
	}
}

func Test_HexAdd(t *testing.T) {
	hexA, _ := NewHex(2, 0)
	hexB, _ := NewHex(0, -1)

	hexC := hexA.Add(*hexB)

	hexD, _ := NewHex(2, -1)
	if hexC != *hexD {
		t.Fatalf(`Hex Addition Failing - (2, 0) + (0, -1) SHOULD = (2, -1)`)
	}
}

func Test_HexSubtract(t *testing.T) {
	hexA, _ := NewHex(2, 1)
	hexB, _ := NewHex(2, -1)

	hexC := hexA.Subtract(*hexB)
	hexR, _ := NewHex(0, 2)

	if hexC != *hexR {
		t.Fatalf(`Hex Subtraction Failing - (2, 1) - (2, -1) SHOULD = (0, 2)`)
	}
}

func Test_HexMultiply(t *testing.T) {
	hexA, _ := NewHex(2, 0)

	hexM := hexA.Multiply(2)
	hexR, _ := NewHex(4, 0)

	if hexM != *hexR {
		t.Fatalf(`Hex Multiplication Failing`)
	}
}

func Test_HexDistance(t *testing.T) {
	hexA, _ := NewHex(2, 2) //s = -4, distance is 4 from 0
	hexB, _ := NewHex(0, 0)
	hexC, _ := NewHex(6, 0) //distance is 4 from hexA, 6 from 0, s = -6
	hexD, _ := NewHex(3, -1) //s = -2, distance is 3 from 0

	testDistance := hexA.Distance(*hexB)
	if testDistance != 4 {
		t.Fatalf(`Hex Distance algorithm failing - distance from (2, 2) -> (0, 0) should be 4, not %v`, testDistance)
	}

	testDistance = hexA.Distance(*hexC)
	if testDistance != 4 {
		t.Fatalf(`Hex Distance algorithm failing - distance from (2, 2) -> (6, 0) should be 4, not %v`, testDistance)
	}

	testDistance = hexB.Distance(*hexD)
	if testDistance != 3 {
		t.Fatalf(`Hex Distance algorithm failure #3 - distance from (0, 0) -> (3, -1) should be 3, not %v`, testDistance)
	}
}

func Test_HexDirection(t *testing.T) {
	hexA, _ := NewHex(0, 0)
	hexSouth := hexA.Neighbor(South) //Should be (0, 1)

	testHex, _ := NewHex(0, 1)
	if hexSouth != *testHex {
		t.Fatalf(`Hex Neighbor failure - new hex should be (0, 1)`)
	}
}