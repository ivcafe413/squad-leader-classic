package handlers

import (
	//"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/room"
	//"github.com/vagrant-technology/squad-leader/session"
)

// ----- Lobby Connection: Active messaging to lobby for user ready status
func LobbyConnection(c *websocket.Conn) {
	//Requires Room and User IDs to function
	username := c.Params("user")
	roomID := c.Params("room")

	r := room.Get(roomID)
	if r == nil {
		//Room Not Found
		fmt.Println("Lobby Connection Error: Room Not Found")
		return //errors.New("room not found") //Call deferred close
	}

	var user *auth.User
	if user = auth.GetUserByName(username); user == nil {
		fmt.Println("Lobby Connection Error: User Not Found: " + username)
		return //errors.New("user not found")
	}

	fmt.Println("Creating Lobby Client for " + user.Username)
	// Create and Register Client to the room Lobby Hub
	client := r.NewClient(c, user)

	// TODO: Write out the initial lobby state on initial connection
	// message, _ := json.Marshal(r.LobbyState())

	// Start read/write client connection
	go client.ConfigureRead()
	go client.ConfigureWrite()
}

func GameConnection(c *websocket.Conn) {
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
}
