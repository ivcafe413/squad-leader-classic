package pieces

import (
	"strings"
	"errors"
)

type GamePiece interface {
	//Get/Set Metadata values
	//Register a Set of Valid Metadata key/value pairs
	//Strategy pattern setters for functionality (first-class functions)
}

type GridPiece struct {
	name string
	//Piece metadata?
}

func NewGridPiece(n string) (*GridPiece, error) {
	if strings.TrimSpace(n) == "" {
		return nil, errors.New("cannot pass empty grid piece name")
	}

	//Piece metadata TODO

	gp := GridPiece { n }
	return &gp, nil
}