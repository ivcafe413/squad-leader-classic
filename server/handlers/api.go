package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/session"
)

// ----- Create Room - RESTful API Call to create room & user
func CreateRoom(c *fiber.Ctx) error {
	//Populate instance of user
	//TODO: Abstract out Auth flow, User Create should not be side effect of Create Room
	u := struct {
		Username string `json:"username"`
	}{}

	// JSON Unmarahal MUST be on pointer struct, NOT non-pointer
	if err := c.BodyParser(&u); err != nil {
		fmt.Println("Create Room Error: " + err.Error())
		return err
	}

	//fmt.Println("Room Owner: " + u.Username)
	user := auth.NewUser(u.Username)
	roomID := session.NewRoom(user)

	return c.JSON(&fiber.Map{
		"success": true,
		"roomID":  roomID,
		"user":    user,
	})
	//return c.SendString(roomID)
}

// ----- Join Room - Websocket connection to open room connection
func JoinRoom(c *fiber.Ctx) error {
	//Requires Room and User IDs to function
	// userID, _ := strconv.Atoi(c.Params("user"))
	joiner := struct {
		RoomID string `json:"roomID"`
		Username string `json:"username"`
	}{}

	//client := new(game.Client)
	if err := c.BodyParser(&joiner); err != nil {
		fmt.Println("Join Room Error: " + err.Error())
		return err
	}

	//roomID := uuid.MustParse(joiner.Room)
	var room *session.Room
	if room = session.GetRoom(joiner.RoomID); room == nil {
		//Room Not Found
		fmt.Println("Join Room Error: Room Not Found")
		return errors.New("room not found")
	}

	var user *auth.User
	if user = auth.GetUserByName(joiner.Username); user == nil {
		fmt.Println("Join Room Error: User Not Found")
		return errors.New("user not found")
	}

	if err := room.Join(user); err != nil {
		fmt.Println("Join Room Error: " + err.Error())
		return err
	}

	return nil
}
