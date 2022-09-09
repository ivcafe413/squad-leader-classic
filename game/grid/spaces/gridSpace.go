package spaces

import (
)
type GridSpace interface {
	Equals(a, b *GridSpace) bool
	Add(a, b *GridSpace) *GridSpace
	Subtract(a, b *GridSpace) *GridSpace
	Multiply(a *GridSpace, b int) *GridSpace
	Distance(a, b *GridSpace) int
	Neighbor(a *GridSpace, d HexDirection) *GridSpace
}
