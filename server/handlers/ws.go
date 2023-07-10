package handlers

import (
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/room"
)

// ----- Lobby Connection: Active messaging to lobby for user ready status
func LobbyConnection(c *websocket.Conn) {
	//Requires Room and User IDs to function
	username := c.Params("user")
	roomID := c.Params("room")

	log.Println(username + " attempting to join Room " + roomID)

	r := room.Get(roomID)
	if r == nil {
		//Room Not Found
		log.Println("Lobby Connection Error: Room Not Found")
		return
	}

	var user *auth.User
	if user = auth.GetUserByName(username); user == nil {
		log.Println("User Not Found: " + username)
		log.Println("Creating User...")
		user = auth.NewUser(username)
	}

	if err := r.Join(user); err != nil {
		log.Println("Join Room Error: " + err.Error())
		return
	}

	log.Println("Creating Lobby Client for " + user.Username)
	// Create and Register Client to the room Lobby Hub
	client := r.NewClient(c, user)

	defer client.Close()

	// Start read/write client connection
	go client.ConfigureWrite()
	client.ConfigureRead()
}

func GameConnection(c *websocket.Conn) {
	username := c.Params("user")
	gameID := c.Params("game")

	log.Println(username, " attempting to join Game ", gameID)

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
}
