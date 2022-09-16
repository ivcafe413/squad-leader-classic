package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
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
	roomID := game.NewRoom(*user)

	return c.JSON(&fiber.Map{
		"success": true,
		"room":    roomID,
		"user":    user,
	})
	//return c.SendString(roomID)
}

// ----- Join Room - Websocket connection to open room connection
func JoinRoom(c *websocket.Conn) {
	defer func() {
		game.RemoveClientConnection <- c
		c.Close()
	}()

	//Requires Room and User IDs to function
	roomID := uuid.MustParse(c.Params("room"))
	userID, _ := strconv.Atoi(c.Params("user"))

	client := new(game.Client)
	if room, exists := game.Rooms[roomID]; !exists {
		//Room Not Found
		fmt.Println("Room Not Found")
		return //Call deferred close
	} else {
		client.Room = room
	}

	if user, exists := store.Users[userID]; !exists {
		fmt.Println("User Not Found")
		return
	} else {
		client.User = user
	}

	cc := new(game.ClientConnection)
	cc.Client = client
	cc.Conn = c
	// Send to Channel
	game.RegisterClientConnection <- cc

	// for {
	// 	//Infinite loop?
	// }

	// Connection over
	// Should fire deferred unregister
}
