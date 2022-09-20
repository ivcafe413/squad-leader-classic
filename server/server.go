package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/game"
	"github.com/vagrant-technology/squad-leader/handlers"
)

func ConfigureWS(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:room/:user", websocket.New(handlers.JoinRoom))

	go game.ClientHub()
}

func ConfigureApi(app *fiber.App) {
	api := app.Group("/api")

	// ----- Web API section -----
	api.Get("/", func(c *fiber.Ctx) error {
		//TODO: Return Homepage/app (UI/static)
		return c.SendString("Hello World")
		//TODO: In Test, use root for stats/debug
	})

	api.Post("/CreateRoom", handlers.CreateRoom)
}

func main() {
	serverConfig := fiber.Config{
		ServerHeader: "Squad Leader Application Server",
	}
	app := fiber.New(serverConfig)

	ConfigureWS(app)
	ConfigureApi(app)
	//apiGroup := app.Group("/api")
	//Router(app)

	app.Listen(":3001")
}
