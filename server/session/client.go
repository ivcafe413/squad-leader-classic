package session

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/auth"
)

type messageReader func([]byte) error

// type Stateful interface {
// 	ReportState() any
// }

type Client[T StateHandler] struct {
	hub        *ClientHub[T]
	connection *websocket.Conn
	user       *auth.User
	reader		messageReader
	writer		chan []byte
	closed bool
}

// Client implementation of websocket Client interface
func (client *Client[T]) ConfigureRead() {
	defer client.Close()

	// Reading incoming from websocket -> client (in)
	for {
		mType, message, err := client.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				//log.Println("read error:", err)
				fmt.Println("read from client error: " + err.Error())
			}

			break // Break out of the for loop
		}

		if mType == websocket.TextMessage {
			// When we receive client messages, we use the Process strategy
			if client.reader(message) != nil {
				fmt.Println("client msg process error: " + err.Error())
				break // out of for loop
			}
		}
	}
}

func (client *Client[T]) ConfigureWrite() {
	defer client.Close()

	// Write messages that end up in client write channel -> websocket (out)
	for {
		message, ok := <- client.writer
		if !ok {
			// Channel closed, end gracefully
			client.connection.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		if err := client.connection.WriteMessage(websocket.TextMessage, message); err != nil {
			// Client Connection write error
			//hub.Remove <- conn
			client.connection.WriteMessage(websocket.CloseMessage, []byte{})
			//conn.Close()
			return
		}
	}
}

func (client *Client[T]) Close() {
	client.hub.Remove <- client.connection //Needs to be idempotent

	if !client.closed {
		client.connection.Close()
		client.closed = true
	}
}

// ----- 

func (client *Client[T]) GetUser() *auth.User {
	return client.user
}

// func (client *Client[T]) Process(msg string) error {
// 	return client.processor(msg)
// }

// func (client *Client[T]) ReadMessage() (int, []byte, error) {
// 	return client.connection.ReadMessage()
// }

func NewClient[T StateHandler](hub *ClientHub[T], conn *websocket.Conn, user *auth.User) *Client[T] {
	client := new(Client[T])

	client.hub = hub
	client.connection = conn
	client.user = user
	client.reader = T.UserInput(user, user, )
	client.writer = make(chan []byte) //256?
	client.closed = false

	// Go ahead and Register the client to the hub, since we have it
	hub.Register <- client

	return client
}
