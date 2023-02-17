package messaging

import (
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/auth"
)

//type messageReader func([]byte) error

// type Stateful interface {
// 	ReportState() any
// }

type Client struct {
	hub        *ClientHub
	connection *websocket.Conn
	user       *auth.User
	//reader		messageReader
	writer chan []byte
	closed bool
}

// Client implementation of websocket Client interface
func (client *Client) ConfigureRead() {
	//defer client.Close()

	// Reading incoming from websocket -> client (in)
	for {
		mType, message, err := client.connection.ReadMessage()
		if err != nil {
			log.Println("client read error: " + err.Error())
			//if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			//fmt.Println(fmt.Sprintf("client read error: %v", err))
			//}
			break // Break out of the for loop, reach deferred functionality
		}

		if mType == websocket.TextMessage {
			// When we receive client messages, we use the Process strategy
			// Pipe the message into the hub's input channel for processing
			log.Println("message from client: " + string(message))
			client.hub.input <- message
		}
	}
}

func (client *Client) ConfigureWrite() {
	//defer client.Close()

	// Write messages that end up in client write channel -> websocket (out)
	for {
		message, ok := <-client.writer
		if !ok {
			// Channel closed, end gracefully
			log.Println("Client write channel closed, closing client...")
			client.connection.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		if err := client.connection.WriteMessage(websocket.TextMessage, message); err != nil {
			//Client Connection write error
			log.Println("client write error: " + err.Error())
			//hub.Remove <- conn
			client.connection.WriteMessage(websocket.CloseMessage, []byte{})
			//conn.Close()
			return
		}
	}
}

func (client *Client) Close() {
	log.Println("Closing client for " + client.GetUser().ID.String() + "...")
	client.hub.Remove <- client.connection //Needs to be idempotent

	if !client.closed {
		client.connection.Close()
		client.closed = true
	}
	log.Println("Client closed")
}

// -----

func (client *Client) GetUser() *auth.User {
	return client.user
}

// func (client *Client[T]) Process(msg string) error {
// 	return client.processor(msg)
// }

// func (client *Client[T]) ReadMessage() (int, []byte, error) {
// 	return client.connection.ReadMessage()
// }

func NewClient(hub *ClientHub, conn *websocket.Conn, user *auth.User) *Client {
	log.Println("Creating new WS client for " + user.ID.String() + "...")
	client := new(Client)

	client.hub = hub
	client.connection = conn
	client.user = user
	//client.reader = T.UserInput(user, user, )
	client.writer = make(chan []byte) //256?
	client.closed = false

	log.Println("New messaging client created!")
	return client
}
