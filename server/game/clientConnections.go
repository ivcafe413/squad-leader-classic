package game

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
)

type ClientHub[T any] struct {
	clients  map[*websocket.Conn]T //Dynamic/open
	register chan Client[T]
}

type Client[T any] struct {
	connection *websocket.Conn
	connEntity T
}

// type ClientConnection struct {
// 	Client *Client
// 	Conn   *websocket.Conn
// }

//var clients = make(map[*websocket.Conn]*Client)

//var lobbyHubs =

// Channels
// var RegisterClientConnection = make(chan *ClientConnection)
// var RemoveClientConnection = make(chan *websocket.Conn)

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

func RunLobbyHub() {

}
