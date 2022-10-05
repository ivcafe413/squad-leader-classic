package room

import (
	//"errors"
	//"fmt"
	"encoding/json"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/session"
)

//type LobbyConnections ClientConnections[*Lobby]

type Lobby struct {
	Users map[*auth.User]bool // Ready Map
	//Room  *Room
	hub *session.ClientHub[*Lobby]
}

// Implement stateful interface
func (lobby *Lobby) MarshalJSON() ([]byte, error) {
	//Return the user lobby into Marshalable for broadcast
	
	jsonLobby := make(map[string]bool)
	for k, v := range lobby.Users {
		jsonLobby[k.Username] = v
	}

	return json.Marshal(jsonLobby)
}

func (lobby *Lobby) Start() {
	lobby.hub = session.NewClientHub(lobby)

	go lobby.hub.Start()
}

// func NewLobby() *Lobby {
// 	lobby := new(Lobby)
// 	//lobby.Users = make(map[*auth.User]bool)
// 	//lobby.Room = r

// 	// Start a messaging Hub for this lobby
// 	lobby.hub = session.NewClientHub(lobby)
// 	go lobby.hub.StartHub()

// 	return lobby
// }

func (lobby *Lobby) Clients() session.ClientConnections[*Lobby] {
	return lobby.hub.Clients
}

func (lobby *Lobby) NewClient(c *websocket.Conn, user *auth.User) *session.Client[*Lobby] {
	return session.NewClient(lobby.hub, c, user)
}

// func (r *Room) RemoveClient(c *websocket.Conn) {
// 	r.lobby.hub.Remove <- c
// }

func (lobby *Lobby) ClientInput(user *auth.User, msg []byte) ([]byte, error) {
	//Make change to room Lobby State/Ready State
	// Expecting raw string in this case // TODO: Implemented different generic expected types
	if string(msg) == "ready" {
		lobby.Users[user] = true
	} else {
		lobby.Users[user] = false
	}

	return json.Marshal(lobby)
}