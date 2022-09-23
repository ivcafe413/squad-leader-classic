package session

import (
	//"errors"

	"github.com/vagrant-technology/squad-leader/auth"
)

type Lobby struct {
	Users map[*auth.User]bool // Ready Map
	//Room  *Room
	hub *ClientHub[*Lobby]
}

// Implement stateful interface
func (lobby *Lobby) ReportState() any {
	//Marshal the user lobby into JSON for broadcast
	//flatLobby, _ := room.MarshalLobby()
	jsonLobby := make(map[string]bool)
	for k, v := range lobby.Users {
		jsonLobby[k.Username] = v
	}

	return jsonLobby
}

func NewLobby() *Lobby {
	lobby := new(Lobby)
	lobby.Users = make(map[*auth.User]bool)
	//lobby.Room = r

	// Start a messaging Hub for this lobby
	lobby.hub = NewClientHub(lobby)
	go lobby.hub.StartHub()

	return lobby
}