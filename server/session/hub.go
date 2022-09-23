package session

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
)

type ClientHub[T Stateful] struct {
	//entity    T
	clients   map[*websocket.Conn]*Client[T] //
	register  chan *Client[T]
	remove    chan *websocket.Conn
	broadcast chan T
}

func (hub *ClientHub[T]) StartHub() {
	for {
		select {
		case rc := <-hub.register:
			hub.clients[rc.Connection] = rc
			fmt.Println("Client connection established")

		case dc := <-hub.remove:
			delete(hub.clients, dc)
			fmt.Println("Client connection closed")

		case state := <-hub.broadcast:
			message, _ := json.Marshal(state.ReportState())
			for conn := range hub.clients {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					// Client Connection write error
					hub.remove <- conn
					conn.WriteMessage(websocket.CloseMessage, []byte{})
					conn.Close()
				}
			}
		}
	}
}

func NewClientHub[T Stateful](v T) *ClientHub[T] {
	hub := new(ClientHub[T])
	//hub.hubEntity = v
	hub.clients = make(map[*websocket.Conn]*Client[T])
	hub.register = make(chan *Client[T])
	hub.remove = make(chan *websocket.Conn)
	hub.broadcast = make(chan T)

	return hub
}
