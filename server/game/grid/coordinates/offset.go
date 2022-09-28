package coordinates

type OffsetCoordinates struct {
	c, r int
}

type Offset int
const (
	OddQ = -1
	EvenQ = 1
)

//TODO: Interface/Embed in a way to properly strategy pattern the 'flat' version
func (hc HexCoordinates) ConvertToOffset(o Offset) OffsetCoordinates {
	c := hc.q
	r := hc.r + int(hc.q + (int(o) * (hc.q & 1)) / 2)

	return OffsetCoordinates{ c, r }
}