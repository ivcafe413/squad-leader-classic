package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vagrant-technology/squad-leader/game"
	"github.com/vagrant-technology/squad-leader/store"
)

// ----- Create Room - RESTful API Call to create room & user
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
	roomID := game.NewRoom(user)

	return c.JSON(&fiber.Map{
		"success": true,
		"room":    roomID,
		"user":    user,
	})
	//return c.SendString(roomID)
}

// ----- Join Room - Websocket connection to open room connection
func JoinRoom(c *fiber.Ctx) error {
	//Requires Room and User IDs to function
	// userID, _ := strconv.Atoi(c.Params("user"))
	joiner := struct {
		Room string `json:"roomID`
		User string `json:"username"`
	}{}

	//client := new(game.Client)
	if err := c.BodyParser(&joiner); err != nil {
		fmt.Println("Join Room Error: " + err.Error())
		return err
	}

	//roomID := uuid.MustParse(joiner.Room)
	var room *game.Room
	if room = game.GetRoom(joiner.Room); room == nil {
		//Room Not Found
		fmt.Println("Join Room Error: Room Not Found")
		return errors.New("Room Not Found")
	}

	var user *store.User
	if user = store.LookupUser(joiner.User); user == nil {
		fmt.Println("Join Room Error: User Not Found")
		return errors.New("User Not Found")
	}

	if err := room.JoinLobby(user); err != nil {
		fmt.Println("Join Room Error: " + err.Error())
		return err
	}

	return nil
}
