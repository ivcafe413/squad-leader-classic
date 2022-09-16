package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/handlers"
)

func Router(app *fiber.App) {
	// ----- Web app section -----
	app.Get("/", func(c *fiber.Ctx) error {
		//TODO: Return Homepage/app (UI/static)
		return c.SendString("Hello World")
		//TODO: In Test, use root for stats/debug
	})

	app.Post("/CreateRoom", handlers.CreateRoom)

	app.Get("/ws/:room/:user", websocket.New(handlers.JoinRoom))
}
