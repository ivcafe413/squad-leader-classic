package game

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/store"
)

type ClientHub[T any] struct {
	hubEntity *T
	clients   map[*websocket.Conn]*Client[T] //
	Register  chan *Client[T]
	Remove    chan *websocket.Conn
	Broadcast chan []byte
}

type Client[T any] struct {
	hub        *ClientHub[T]
	connection *websocket.Conn
	user       *store.User
}

func NewClient[T any](hub *ClientHub[T], conn *websocket.Conn, user *store.User) *Client[T] {
	client := new(Client[T])

	client.hub = hub
	client.connection = conn
	client.user = user

	return client
}

func NewClientHub[T any](v *T) *ClientHub[T] {
	hub := new(ClientHub[T])
	hub.hubEntity = v
	hub.clients = make(map[*websocket.Conn]*Client[T])
	hub.Register = make(chan *Client[T])
	hub.Remove = make(chan *websocket.Conn)
	hub.Broadcast = make(chan []byte)

	return hub
}

func (hub *ClientHub[T]) StartHub() {
	for {
		select {
		case rc := <-hub.Register:
			hub.clients[rc.connection] = rc
			fmt.Println("Client connection established")

		case dc := <-hub.Remove:
			delete(hub.clients, dc)
			fmt.Println("Client connection closed")

		case message := <-hub.Broadcast:
			for conn := range hub.clients {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					// Client Connection write error
					hub.Remove <- conn
					conn.WriteMessage(websocket.CloseMessage, []byte{})
					conn.Close()
				}
			}
		}
	}
}
