package room

import (
	"encoding/json"
	"errors"
	"log"

	//"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/messaging"
)

var rooms = make(map[uuid.UUID]*Room)

// type GameRoom interface {
// 	Join(*auth.User)
// 	Close()
// }

type Room struct {
	ID       uuid.UUID           `json:"id"`
	Owner    *auth.User          `json:"owner"`
	Users    map[*auth.User]bool // Ready Map
	hub      *messaging.ClientHub
	IsClosed bool `json:"isClosed"`
}

func (r *Room) Join(user *auth.User) error {
	// If the User is not already in the Lobby
	if _, exists := r.Users[user]; exists {
		return errors.New("user already in lobby")
	}

	r.Users[user] = false
	//readyMsg, _ := r.ClientInput(user, []byte("not ready"))

	//Broadcast state change to clients
	//r.hub.Broadcast <- readyMsg
	return nil
}

func (r *Room) NewClient(c *websocket.Conn, user *auth.User) *messaging.Client {
	nc := messaging.NewClient(r.hub, c, user)
	r.hub.Register <- nc

	return nc
}

func (r *Room) Close() error {
	r.hub.Stop()

	r.IsClosed = true
	return nil
}

func (r *Room) ClientInput(user *auth.User, msg []byte) ([]byte, error) {
	//Make change to room Lobby State/Ready State
	// Expecting raw string in this case // TODO: Implemented different generic expected types
	if string(msg) == "ready" {
		r.Users[user] = true
	} else {
		r.Users[user] = false
	}

	return r.MarshalJSON()
}

func (r *Room) MarshalJSON() ([]byte, error) {
	//Return the user lobby into Marshalable for broadcast

	roomJson := make(map[string]bool)
	for k, v := range r.Users {
		roomJson[k.Username] = v
	}

	return json.Marshal(roomJson)
}

// ---------- Static Methods ----------

func NewRoom(user *auth.User) string {
	log.Println("Creating new room...")
	roomID := uuid.New()

	room := new(Room)
	room.ID = roomID
	room.Owner = user
	room.IsClosed = false

	room.Users = make(map[*auth.User]bool)
	//room.Users[user] = false

	rooms[roomID] = room
	room.hub = messaging.NewClientHub()
	go room.hub.Start()

	//readyMsg, _ := room.ClientInput(user, []byte("not ready"))

	//Broadcast state change to clients
	//room.hub.Broadcast <- readyMsg

	log.Println("Room " + room.ID.String() + " created")
	return room.ID.String()
}

func Get(room string) *Room {
	roomID := uuid.MustParse(room)
	return rooms[roomID]
}
