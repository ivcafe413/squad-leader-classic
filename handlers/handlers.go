package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vagrant-technology/squad-leader/game"
	"github.com/vagrant-technology/squad-leader/store"
)

func CreateRoom(c *fiber.Ctx) error {
	//Populate instance of user
	user := new(store.User)
	if err := c.BodyParser(user); err != nil {
		// c.Status(500).JSON(&fiber.Map{
		// 	"success": false,
		// 	"error":   err,
		// })
		return err
	}

	roomID := game.NewRoom()
	return c.JSON(&fiber.Map{
		"success": true,
		"room":    roomID,
	})
}
