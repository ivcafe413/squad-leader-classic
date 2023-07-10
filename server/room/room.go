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

type Room struct {
	ID       uuid.UUID           `json:"id"`
	Owner    *auth.User          `json:"owner"`
	Users    map[*auth.User]bool // Ready Map
	hub      *messaging.ClientHub
	IsClosed bool `json:"isClosed"`
}

type UserReady struct {
	User  *auth.User `json:"user"`
	Ready bool       `json:"ready"`
}

func (r *Room) Join(user *auth.User) error {
	// If the User is not already in the Lobby
	if _, exists := r.Users[user]; exists {
		return errors.New("user already in lobby")
	}

	r.Users[user] = false

	return nil
}

func (r *Room) NewClient(c *websocket.Conn, user *auth.User) *messaging.Client {
	reader := func(msg []byte) interface{} {
		stringMsg := string(msg)
		//log.Println("checking incoming client message: ", stringMsg)
		readyState := stringMsg == "ready"
		//log.Println("checking translated ready state: ", readyState)
		rawMsg := UserReady{
			User:  user,
			Ready: readyState,
		}

		return rawMsg
	}

	nc := messaging.NewClient(r.hub, c, user, reader)
	r.hub.Register <- nc

	//Broadcast state change to clients
	msg, _ := r.MarshalJSON()
	r.hub.Broadcast <- msg

	return nc
}

func (r *Room) Close() error {
	r.hub.Stop()

	r.IsClosed = true
	return nil
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
	rooms[roomID] = room

	//Start the Message Hub w/Input Processor
	processor := func(msg interface{}) []byte {
		switch x := msg.(type) {
		case UserReady:
			room.Users[x.User] = x.Ready

			readyMsg, _ := room.MarshalJSON()
			//log.Println("marshalled room state: ", string(readyMsg))
			return readyMsg

		default:
			return nil
		}
	}
	room.hub = messaging.NewClientHub(processor)
	go room.hub.Start()

	log.Println("Room " + room.ID.String() + " created")
	return room.ID.String()
}

func Get(room string) *Room {
	roomID := uuid.MustParse(room)
	return rooms[roomID]
}
