package session

import (
	//"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
)

// Type Alias for Connection mapping
type ClientConnections[T StateHandler] map[*websocket.Conn]*Client[T]

type ClientHub[T StateHandler] struct {
	//entity    T
	Clients   ClientConnections[T] //
	Register  chan *Client[T]
	Remove    chan *websocket.Conn
	Broadcast chan interface{}
	done 	  chan bool
}

func (hub *ClientHub[T]) Start() {
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

		case message := <-hub.Broadcast:
			//message, _ := json.Marshal(state.ReportState())
			for conn := range hub.Clients {
				if err := conn.WriteJSON(message); err != nil {
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

func (hub *ClientHub[T]) Stop() {
	<-hub.done // Signal the Hub closed, stops the hub goroutine
	for _, client := range hub.Clients {
		client.Close()
	}

	close(hub.Register)
	close(hub.Remove)
	close(hub.Broadcast)
}

func NewClientHub[T StateHandler](v T) *ClientHub[T] {
	hub := new(ClientHub[T])
	
	hub.Clients = make(map[*websocket.Conn]*Client[T])
	hub.Register = make(chan *Client[T])
	hub.Remove = make(chan *websocket.Conn)
	hub.Broadcast = make(chan interface{})

	return hub
}
