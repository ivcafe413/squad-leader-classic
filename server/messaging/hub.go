package messaging

import (
	//"encoding/json"

	"log"

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

	done  chan bool
	input chan []byte
}

func (hub *ClientHub) Start() {
	// Goroutine
	log.Println("Client messaging hub has started")
	for {
		select {
		case rc := <-hub.Register:
			hub.Clients[rc.connection] = rc // Upsert-ish
			//fmt.Println("Client connection established")
			log.Println("Client connection established - ", rc.GetUser().ID.String())

		case dc := <-hub.Remove:
			if deleted, ok := hub.Clients[dc]; ok {
				delete(hub.Clients, dc)
				log.Println("Client connection removed - ", deleted.GetUser().ID.String())
			} else {
				// Ignore if not existing, prevent panic
				log.Println("Client conn not found on attempt to delete")
			}

		case message := <-hub.Broadcast:
			for conn := range hub.Clients {
				if err := conn.WriteJSON(message); err != nil {
					// Client Connection write error
					log.Println("Messaging hub write error - " + err.Error())
					//fmt.Println(fmt.Sprintf("hub broadcast error: %v", err))

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
	log.Println("Stopping Client messaging hub...")

	for _, client := range hub.Clients {
		client.Close()
	}

	<-hub.done // Signal the Hub closed, stops the hub goroutine

	close(hub.Register)
	close(hub.Remove)
	close(hub.Broadcast)
	log.Println("Client messaging hub has closed")
}

// ----- ----- Static functions ----- -----

func NewClientHub() *ClientHub {
	log.Println("Creating new messaging hub for ws...")
	hub := new(ClientHub)

	hub.Clients = make(map[*websocket.Conn]*Client)
	hub.Register = make(chan *Client)
	hub.Remove = make(chan *websocket.Conn)
	hub.Broadcast = make(chan interface{})

	hub.done = make(chan bool, 1)
	hub.input = make(chan []byte)

	log.Println("New messaging hub created")
	return hub
}
