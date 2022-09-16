package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/game"
)

func ConfigureWS(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
}

func main() {
	serverConfig := fiber.Config{
		ServerHeader: "Squad Leader Application Server",
	}
	app := fiber.New(serverConfig)

	ConfigureWS(app)

	go game.ClientHub()

	Router(app)

	app.Listen(":3000")
}
