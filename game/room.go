package game

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/game/grid"
	"github.com/vagrant-technology/squad-leader/store"
)

var Rooms = make(map[uuid.UUID]*Room)

type Room struct {
	ID uuid.UUID `json:"id"`
	Owner store.User `json:"owner"`
	Grid *grid.HexGrid
}

func NewRoom(user store.User) string {
	roomID := uuid.New()
	fmt.Println("Test New UUID: " + roomID.String())
	// newRoom := Room {
	// 	ID: gameID,
	// 	grid: *grid.NewHexGrid(33, 10),
	// }
	newRoom := new(Room)
	newRoom.ID = roomID
	newRoom.Owner = user
	newRoom.Grid = grid.NewHexGrid(33, 10)

	Rooms[roomID] = newRoom

	return newRoom.ID.String()
}
