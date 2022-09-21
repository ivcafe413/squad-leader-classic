package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/game"
	"github.com/vagrant-technology/squad-leader/store"
)

// ----- Lobby Connection: Active messaging to lobby for user ready status
func LobbyConnection(c *websocket.Conn) {
	// Close WS connections appropriately
	defer func() {
		//game.RemoveClientConnection <- c
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

func GameConnection(c *websocket.Conn) {
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
