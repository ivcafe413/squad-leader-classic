package room

import (
	"errors"
	//"fmt"

	//"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/auth"
	//"github.com/vagrant-technology/squad-leader/session"
)

var rooms = make(map[uuid.UUID]*Room)

type GameRoom interface {
	Join(*auth.User)
	Close()
}

type Room struct {
	ID    uuid.UUID		`json:"id"`
	Owner *auth.User	`json:"owner"`
	*Lobby
	IsClosed bool 		`json:"isClosed"`
}

func NewRoom(user *auth.User) string {
	roomID := uuid.New()

	room := new(Room)
	room.ID = roomID
	room.Owner = user
	room.IsClosed = false
	
	//room.lobby = NewLobby()

	rooms[roomID] = room
	room.Start()

	return room.ID.String()
}

func Get(room string) *Room {
	roomID := uuid.MustParse(room)
	return rooms[roomID]
}

// func GetLobby(r *Room) *Lobby {
// 	return r.lobby
// }

// ----- TODO: Interface?-----

func (r *Room) Join(user *auth.User) error {
	// If the User is not already in the Lobby
	if _, exists := r.Users[user]; exists {
		return errors.New("user already in lobby")
	}

	//r.Users[user] = false
	readyMsg, _ := r.ClientInput(user, []byte("not ready"))

	//Broadcast state change to clients
	r.hub.Broadcast <- readyMsg
	return nil
}

func (r *Room) Close() error {
	r.hub.Stop()

	r.IsClosed = true
	return nil
}
