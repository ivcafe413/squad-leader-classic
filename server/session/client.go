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
	Connection *websocket.Conn
	user       *auth.User
	processor  messageProcessor
}

func (client *Client[T]) Process(msg string) error {
	return client.processor(msg)
}

func (client *Client[T]) Close() {
	client.hub.remove <- client.Connection
	client.Connection.Close()
}

func NewClient[T Stateful](hub *ClientHub[T], conn *websocket.Conn, user *auth.User, processor messageProcessor) *Client[T] {
	client := new(Client[T])

	client.hub = hub
	client.Connection = conn
	client.user = user
	client.processor = processor

	// Go ahead and Register the client to the hub, since we have it
	hub.register <- client

	return client
}
