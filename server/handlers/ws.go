package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/game"
	"github.com/vagrant-technology/squad-leader/store"
)

// ----- Lobby Connection: Active messaging to lobby for user ready status
func LobbyConnection(c *websocket.Conn) error {
	var room *game.Room // Need Lobby Hub early
	room = game.GetRoom(c.Params("room"))
	if room == nil {
		//Room Not Found
		fmt.Println("Lobby Connection Error: Room Not Found")
		return errors.New("room not found") //Call deferred close
	}
	hub := room.Lobby.Hub

	// Close WS connections appropriately
	defer func() {
		hub.Remove <- c
		c.Close()
	}()

	//Requires Room and User IDs to function
	//roomID := uuid.MustParse(c.Params("room"))
	//roomID := c.Params("room")
	//userID, _ := strconv.Atoi(c.Params("user"))
	username := c.Params("user")

	//client := new(game.Client)

	var user *store.User
	if user = store.LookupUser(username); user == nil {
		fmt.Println("Lobby Connection Error: User Not Found")
		return errors.New("user not found")
	}

	//cc := new(game.ClientConnection)

	client := game.NewClient(hub, c, user)
	//cc.Client = client
	//cc.Conn = c

	// Register Client to the Lobby Hub
	//game.RegisterClientConnection <- cc
	hub.Register <- client

	//
	for {
		mType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				//log.Println("read error:", err)
				fmt.Println("read from client error: " + err.Error())
			}

			return err
		}

		//if mType == websocket.TextMessage
		//TODO:
	}

	// Connection over
	// Should fire deferred unregister
}

// func GameConnection(c *websocket.Conn) {
// 	defer func() {
// 		game.RemoveClientConnection <- c
// 		c.Close()
// 	}()

// 	//Requires Room and User IDs to function
// 	roomID := uuid.MustParse(c.Params("room"))
// 	userID, _ := strconv.Atoi(c.Params("user"))

// 	client := new(game.Client)
// 	if room, exists := game.Rooms[roomID]; !exists {
// 		//Room Not Found
// 		fmt.Println("Room Not Found")
// 		return //Call deferred close
// 	} else {
// 		client.Room = room
// 	}

// 	if user, exists := store.Users[userID]; !exists {
// 		fmt.Println("User Not Found")
// 		return
// 	} else {
// 		client.User = user
// 	}

// 	cc := new(game.ClientConnection)
// 	cc.Client = client
// 	cc.Conn = c
// 	// Send to Channel
// 	game.RegisterClientConnection <- cc

// 	// for {
// 	// 	//Infinite loop?
// 	// }

// 	// Connection over
// 	// Should fire deferred unregister
// }
