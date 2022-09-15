package main

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// ----- Web app section -----
	app.Get("/", func(c *fiber.Ctx) error {
		//TODO: Return Homepage/app (UI/static)
		return c.SendString("Hello World")
		//TODO: In Test, use root for stats/debug
		
	})

	app.Post("/CreateRoom", func(c *fiber.Ctx) error {
		//Populate instance of user
		user := new(User);
		if err := c.BodyParser(user); err != nil {
			return err
		}
		

		roomID := game.NewRoom()
		return c.SendString(roomID)
	})
}