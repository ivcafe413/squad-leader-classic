package session

import (
	"errors"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/auth"
)

var rooms = make(map[uuid.UUID]*Room)

//type LobbyUsers map[*store.User]bool

type Room struct {
	ID    uuid.UUID   `json:"id"`
	Owner *auth.User `json:"owner"`
	//Lobby map[*store.User]bool //True indicates Ready player
	lobby *Lobby
	//Grid  *grid.HexGrid
}

func NewRoom(user *auth.User) string {
	roomID := uuid.New()

	room := new(Room)
	room.ID = roomID
	room.Owner = user
	//newRoom.Lobby = make(map[*store.User]bool)
	room.lobby = NewLobby()
	// newRoom.Grid = grid.NewHexGrid(33, 10)

	rooms[roomID] = room

	return room.ID.String()
}

func GetRoom(room string) *Room {
	roomID := uuid.MustParse(room)
	return rooms[roomID]
}

// ----- TODO: Interface?-----

func (r *Room) Join(user *auth.User) error {
	// If the User is not already in the Lobby
	if _, exists := r.lobby.Users[user]; exists {
		return errors.New("user already in lobby")
	}

	r.lobby.Users[user] = false

	//Broadcast state change to clients
	r.lobby.hub.broadcast <- r.lobby
	return nil
}

func (r *Room) NewLobbySession(c *websocket.Conn, user *auth.User) *Client[*Lobby] {
	// Define/pass message processor strategy
	var lobbyMessage = func(msg string) error {
		//Make change to room Lobby State/Ready State
		if msg == "ready" {
			r.lobby.Users[user] = true
		} else {
			r.lobby.Users[user] = false
		}

		r.lobby.hub.broadcast <- r.lobby
		return nil
	}

	return NewClient(r.lobby.hub, c, user, lobbyMessage)
}

func (r *Room) LobbyState() any {
	return r.lobby.ReportState()
}
