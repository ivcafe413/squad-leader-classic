package messaging

import (
	//"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
)

// Type Alias for Connection mapping
type ClientConnections map[*websocket.Conn]*Client

type ClientHub struct {
	//entity    T
	Clients   ClientConnections //
	Register  chan *Client
	Remove    chan *websocket.Conn
	Broadcast chan interface{}

	done 	  chan bool
	input chan []byte
}

func (hub *ClientHub) Start() {
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

func (hub *ClientHub) Stop() {
	<-hub.done // Signal the Hub closed, stops the hub goroutine
	for _, client := range hub.Clients {
		client.Close()
	}

	close(hub.Register)
	close(hub.Remove)
	close(hub.Broadcast)
}

func NewClientHub() *ClientHub {
	hub := new(ClientHub)
	
	hub.Clients = make(map[*websocket.Conn]*Client)
	hub.Register = make(chan *Client)
	hub.Remove = make(chan *websocket.Conn)
	hub.Broadcast = make(chan interface{})
	
	hub.done = make(chan bool, 1)
	hub.input = make(chan []byte)
	

	return hub
}
