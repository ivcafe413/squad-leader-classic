package game

import (
	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/game/grid"
)

type Room struct {
	ID uuid.UUID
	grid grid.HexGrid
}

func NewRoom() string {
	gameID := uuid.New()
	newRoom := Room {
		ID: gameID,
		grid: *grid.NewHexGrid(33, 10),
	}

	return newRoom.ID.String()
}
