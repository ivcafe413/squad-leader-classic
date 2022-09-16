package game

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/store"
)

type Client struct {
	Room *Room
	User *store.User
}

type ClientConnection struct {
	Client *Client
	Conn *websocket.Conn
}

var clients = make(map[*websocket.Conn]*Client)

// Channels
var RegisterClientConnection = make(chan *ClientConnection)
var RemoveClientConnection = make(chan *websocket.Conn)

func ClientHub() {
	for {
		select {
		case cc := <-RegisterClientConnection:
			clients[cc.Conn] = cc.Client
			fmt.Println("Client connection established")

		case c := <-RemoveClientConnection:
			delete(clients, c)
			fmt.Println("Client connection closed")
		}
	}
}