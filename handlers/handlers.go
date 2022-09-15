package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/vagrant-technology/squad-leader/game"
	"github.com/vagrant-technology/squad-leader/store"
)

func CreateRoom(c *fiber.Ctx) error {
	//Populate instance of user
	//user := new(store.User)
	u := struct {
		Username string `json:"username"`
	}{}
	
	// JSON Unmarahal MUST be on pointer struct, NOT non-pointer
	if err := c.BodyParser(&u); err != nil {
		fmt.Println("Create Room Error: " + err.Error())
		return err
	}

	fmt.Println("Room Owner: " + u.Username)
	user := store.NewUser(u.Username)
	roomID := game.NewRoom(*user)

	return c.JSON(&fiber.Map{
		"success": true,
		"room":    roomID,
		"user": user,
	})
	//return c.SendString(roomID)
}
