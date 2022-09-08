package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/game"
)

type GameRoom struct {
	ID   uuid.UUID
	Game game.GameState
}

func CreateRoom() string {
	roomID := uuid.New()
	room := GameRoom{
		ID: roomID,
	}

	return room.ID.String()
}

func main() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/", func(c *fiber.Ctx) error {
		//return c.SendString("Hello World")
		roomID := CreateRoom()
		return c.SendString(roomID)
	})

	app.Listen(":3000")
}
