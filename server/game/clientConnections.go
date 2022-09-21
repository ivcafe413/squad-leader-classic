package game

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/store"
)

type ClientHub[T any] struct {
	hubEntity T
	clients   map[*websocket.Conn]*Client[T] //
	register  chan *Client[T]
	remove    chan *Client[T]
	broadcast chan []byte
}

type Client[T any] struct {
	hub        *ClientHub[T]
	connection *websocket.Conn
	user       *store.User
}

func NewClientHub[T any](v T) *ClientHub[T] {
	hub := new(ClientHub[T])
	hub.hubEntity = v
	hub.clients = make(map[*websocket.Conn]*Client[T])
	hub.register = make(chan *Client[T])
	hub.remove = make(chan *Client[T])
	hub.broadcast = make(chan []byte)

	return hub
}

func (hub *ClientHub[T]) StartHub() {
	for {
		select {
		case rc := <-hub.register:
			hub.clients[rc.connection] = rc
			fmt.Println("Client connection established")

		case dc := <-hub.remove:
			delete(hub.clients, dc.connection)
			fmt.Println("Client connection closed")

		case message := <-hub.broadcast:
			for conn := range hub.clients {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					// Client Connection write error
					hub.remove <- hub.clients[conn]
					conn.WriteMessage(websocket.CloseMessage, []byte{})
					conn.Close()
				}
			}
		}
	}
}
