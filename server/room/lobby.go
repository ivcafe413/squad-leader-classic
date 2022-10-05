package room

import (
	//"errors"

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
func (lobby *Lobby) ReportState() any {
	//Return the user lobby into Marshalable for broadcast
	
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
	lobby.hub = session.NewClientHub(lobby)
	go lobby.hub.StartHub()

	return lobby
}