package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/room"
	"github.com/vagrant-technology/squad-leader/session"
)

// ----- Local WS Connection logic for conn read/writes
func startRead[T session.Stateful](c *session.Client[T]) {
	defer c.Close()

	// Reading incoming from end client -> websocket
	for {
		mType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				//log.Println("read error:", err)
				fmt.Println("read from client error: " + err.Error())
			}

			break // Break out of the for loop
		}

		if mType == websocket.TextMessage {
			// When we receive client messages, we use the Process strategy
			if c.Process(string(message)) != nil {
				fmt.Println("client msg process error: " + err.Error())
				break // out of for loop
			}
		}
	}

	// Done, unblock the main connection thread
}

// ----- Lobby Connection: Active messaging to lobby for user ready status
func LobbyConnection(c *websocket.Conn) {
	//Requires Room and User IDs to function
	username := c.Params("user")
	roomID := c.Params("room")

	//var room *game.Room
	// Need Lobby Hub early
	r := room.GetRoom(roomID)
	if r == nil {
		//Room Not Found
		fmt.Println("Lobby Connection Error: Room Not Found")
		return //errors.New("room not found") //Call deferred close
	}
	// hub := room.Lobby.Hub

	var user *auth.User
	if user = auth.GetUserByName(username); user == nil {
		fmt.Println("Lobby Connection Error: User Not Found: " + username)
		return //errors.New("user not found")
	}

	fmt.Println("Creating Lobby Client for " + user.Username)
	// Create and Register Client to the Lobby Hub
	client := r.NewClient(c, user)

	// Write out the initial lobby state on initial connection
	message, _ := json.Marshal(r.LobbyState())
	if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
		// Client Connection write error
		//hub.Remove <- c
		c.WriteMessage(websocket.CloseMessage, []byte{})
		//c.Close()
		return
	}

	// Start reading from client connection
	go startRead(client)
	// startRead will handle connection close, can exit func gracefully
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
