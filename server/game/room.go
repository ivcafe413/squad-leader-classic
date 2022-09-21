package game

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/game/grid"
	"github.com/vagrant-technology/squad-leader/store"
)

var rooms = make(map[uuid.UUID]*Room)

type LobbyUsers map[*store.User]bool

type Lobby struct {
	Users LobbyUsers
	Room  *Room
}

type Room struct {
	ID    uuid.UUID   `json:"id"`
	Owner *store.User `json:"owner"`
	//Lobby map[*store.User]bool //True indicates Ready player
	Lobby *Lobby
	Grid  *grid.HexGrid
}

func NewRoom(user *store.User) string {
	roomID := uuid.New()
	fmt.Println("Test New UUID: " + roomID.String())
	// newRoom := Room {
	// 	ID: gameID,
	// 	grid: *grid.NewHexGrid(33, 10),
	// }
	room := new(Room)
	room.ID = roomID
	room.Owner = user
	//newRoom.Lobby = make(map[*store.User]bool)
	room.Lobby = room.NewLobby()
	// newRoom.Grid = grid.NewHexGrid(33, 10)

	rooms[roomID] = room

	return room.ID.String()
}

func GetRoom(room string) *Room {
	roomID := uuid.MustParse(room)
	return rooms[roomID]
}

func (r *Room) NewLobby() *Lobby {
	lobby := new(Lobby)
	lobby.Users = make(LobbyUsers)
	lobby.Room = r

	// Start a messaging Hub for this lobby

	return lobby
}

func (r *Room) JoinLobby(user *store.User) error {
	// If the User is not already in the Lobby
	if _, exists := r.Lobby.Users[user]; exists {
		return errors.New("User already in Lobby")
	}

	r.Lobby.Users[user] = false
	return nil
}
