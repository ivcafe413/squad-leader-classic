package session

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
)

// Type Alias for Connection mapping
type ClientConnections[T Stateful] map[*websocket.Conn]*Client[T]

type ClientHub[T Stateful] struct {
	//entity    T
	Clients   ClientConnections[T] //
	Register  chan *Client[T]
	Remove    chan *websocket.Conn
	Broadcast chan T
	done 	  chan bool
}

func (hub *ClientHub[T]) StartHub() {
	// Goroutine
	for {
		select {
		case rc := <-hub.Register:
			hub.Clients[rc.connection] = rc // Upsert-ish
			fmt.Println("Client connection established")

		case dc := <-hub.Remove:
			if _, ok := hub.Clients[dc]; ok {
				delete(hub.Clients, dc)
				fmt.Println("Client connection removed")
			}
			// Ignore if not existing, prevent panic

		case state := <-hub.Broadcast:
			message, _ := json.Marshal(state.ReportState())
			for conn := range hub.Clients {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					// Client Connection write error
					hub.Remove <- conn
					conn.WriteMessage(websocket.CloseMessage, []byte{})
					conn.Close()
				}
			}

		case <-hub.done:
			return // Exit/kill goroutine when hub is done running
		}
	}
}

func (hub *ClientHub[T]) StopHub() {
	<-hub.done // Signal the Hub closed, stops the hub goroutine
	for _, client := range hub.Clients {
		client.Close()
	}

	close(hub.Register)
	close(hub.Remove)
	close(hub.Broadcast)
}

func NewClientHub[T Stateful](v T) *ClientHub[T] {
	hub := new(ClientHub[T])
	
	hub.Clients = make(map[*websocket.Conn]*Client[T])
	hub.Register = make(chan *Client[T])
	hub.Remove = make(chan *websocket.Conn)
	hub.Broadcast = make(chan T)

	return hub
}
