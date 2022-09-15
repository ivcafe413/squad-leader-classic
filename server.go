package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	// "github.com/google/uuid"
	// "github.com/vagrant-technology/squad-leader/game"
	// "github.com/vagrant-technology/squad-leader/store"
)

type client struct {
	
}

func main() {
	serverConfig := fiber.Config {
		ServerHeader: "Squad Leader Application Server",
	}
	app := fiber.New(serverConfig)

	Router(app)

	// ----- WebSockets -----
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Listen(":3000")
}
