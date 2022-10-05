package session

import (
	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/auth"
)

type messageProcessor func(string) error

type Stateful interface {
	ReportState() any
}

type Client[T Stateful] struct {
	hub        *ClientHub[T]
	connection *websocket.Conn
	user       *auth.User
	processor  messageProcessor
	closed bool
}

func (client *Client[T]) GetUser() *auth.User {
	return client.user
}

func (client *Client[T]) Process(msg string) error {
	return client.processor(msg)
}

func (client *Client[T]) ReadMessage() (int, []byte, error) {
	return client.connection.ReadMessage()
}

func (client *Client[T]) Close() {
	client.hub.Remove <- client.connection //Needs to be idempotent

	if !client.closed {
		client.connection.Close()
		client.closed = true
	}
}

func NewClient[T Stateful](hub *ClientHub[T], conn *websocket.Conn, user *auth.User, processor messageProcessor) *Client[T] {
	client := new(Client[T])

	client.hub = hub
	client.connection = conn
	client.user = user
	client.processor = processor
	client.closed = false

	// Go ahead and Register the client to the hub, since we have it
	hub.Register <- client

	return client
}
