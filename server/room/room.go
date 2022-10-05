package room

import (
	"errors"
	"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/session"
)

var rooms = make(map[uuid.UUID]*Room)

//type LobbyUsers map[*store.User]bool

type Room struct {
	ID    uuid.UUID		`json:"id"`
	Owner *auth.User	`json:"owner"`
	lobby *Lobby		`json:"-"`
	//Grid  *grid.HexGrid
	//game *game.Game		`json:"-"`
	IsClosed bool 		`json:"isClosed"`
}

func NewRoom(user *auth.User) string {
	roomID := uuid.New()

	room := new(Room)
	room.ID = roomID
	room.Owner = user
	room.IsClosed = false
	
	room.lobby = NewLobby()

	rooms[roomID] = room

	return room.ID.String()
}

func GetRoom(room string) *Room {
	roomID := uuid.MustParse(room)
	return rooms[roomID]
}

// func GetLobby(r *Room) *Lobby {
// 	return r.lobby
// }

// ----- TODO: Interface?-----

func (r *Room) Join(user *auth.User) error {
	// If the User is not already in the Lobby
	if _, exists := r.lobby.Users[user]; exists {
		return errors.New("user already in lobby")
	}

	r.lobby.Users[user] = false

	//Broadcast state change to clients
	r.lobby.hub.Broadcast <- r.lobby
	return nil
}

func (r *Room) Close() error {
	r.lobby.hub.StopHub()

	r.IsClosed = true
	return nil
}

func (r *Room) GetClients() session.ClientConnections[*Lobby] {
	return r.lobby.hub.Clients
}

func (r *Room) NewClient(c *websocket.Conn, user *auth.User) *session.Client[*Lobby] {
	// Define/pass message processor strategy
	var lobbyMessage = func(msg string) error {
		//Make change to room Lobby State/Ready State
		fmt.Println("Updating Lobby... ")
		fmt.Println("Room: " + r.ID.String())
		fmt.Println("User: " + user.Username)
		fmt.Println("Message: " + msg)

		if msg == "ready" {
			r.lobby.Users[user] = true
		} else {
			r.lobby.Users[user] = false
		}

		r.lobby.hub.Broadcast <- r.lobby
		return nil
	}

	return session.NewClient(r.lobby.hub, c, user, lobbyMessage)
}

// func (r *Room) RemoveClient(c *websocket.Conn) {
// 	r.lobby.hub.Remove <- c
// }

func (r *Room) LobbyState() any {
	return r.lobby.ReportState()
}
