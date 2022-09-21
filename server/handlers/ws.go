package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/game"
	"github.com/vagrant-technology/squad-leader/store"
)

// ----- Local WS Connection logic for conn read/writes
func startRead[T any](c *game.Client[T], cSignal chan bool) {
	// Reading incoming from end client -> websocket
	for {
		mType, message, err := c.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				//log.Println("read error:", err)
				fmt.Println("read from client error: " + err.Error())
			}

			break // Break out of the for loop
		}

		if mType == websocket.TextMessage {
			//TODO: Interface/Strategy pattern
			// if _, isLobby := any(c.).(game.Lobby); isLobby {
			if c.Process(string(message)) != nil {
				fmt.Println("client msg process error: " + err.Error())
				break // out of for loop
			}
		}
	}

	//Done, unblock the main connection thread
	cSignal <- true
}

// ----- Lobby Connection: Active messaging to lobby for user ready status
func LobbyConnection(c *websocket.Conn) {
	signalClose := make(chan bool)

	// Close WS connections appropriately
	defer func() {
		c.Close()
		close(signalClose)
	}()

	//Requires Room and User IDs to function
	username := c.Params("user")
	roomID := c.Params("room")

	//var room *game.Room
	// Need Lobby Hub early
	room := game.GetRoom(roomID)
	if room == nil {
		//Room Not Found
		fmt.Println("Lobby Connection Error: Room Not Found")
		return //errors.New("room not found") //Call deferred close
	}
	hub := room.Lobby.Hub

	defer func() {
		hub.Remove <- c
	}()

	var user *store.User
	if user = store.LookupUser(username); user == nil {
		fmt.Println("Lobby Connection Error: User Not Found")
		return //errors.New("user not found")
	}

	// Define/pass message processor strategy
	var lobbyMessage = func(msg string) error {
		//Make change to room Lobby State/Ready State
		if msg == "ready" {
			room.Lobby.Users[user] = true
		} else {
			room.Lobby.Users[user] = false
		}

		//Marshal the user lobby into JSON for broadcast
		flatLobby, _ := room.MarshalLobby()
		lobMsg, err := json.Marshal(flatLobby)
		if err != nil {
			return err
		}

		hub.Broadcast <- lobMsg
		return nil
	}
	client := game.NewClient(hub, c, user, lobbyMessage)

	// Register Client to the Lobby Hub -- Moved into ClientConn logic above (NewClient)
	//hub.Register <- client

	// Start reading from client connection
	go startRead(client, signalClose)
	// Wait for the read goroutine to be done
	<-signalClose

	// Connection over
	// Should fire deferred closes/unregisters
	//return
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
