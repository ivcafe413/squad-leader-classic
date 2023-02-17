package handlers

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/room"
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
		//fmt.Println("Create Room Error: " + err.Error())
		log.Println("Create room error: " + err.Error())
		return err
	}

	user := auth.NewUser(u.Username)
	roomID := room.NewRoom(user)
	log.Println(user.Username + " successfully created room " + roomID)
	return c.JSON(&fiber.Map{
		"success": true,
		"roomID":  roomID,
		"user":    user,
	})
}

// ----- Join Room - Websocket connection to open room connection
func JoinRoom(c *fiber.Ctx) error {
	//Requires Room and User IDs to function
	// userID, _ := strconv.Atoi(c.Params("user"))
	joiner := struct {
		RoomID   string `json:"roomID"`
		Username string `json:"username"`
	}{}

	//client := new(game.Client)
	if err := c.BodyParser(&joiner); err != nil {
		fmt.Println("Join Room Error: " + err.Error())
		return err
	}

	//roomID := uuid.MustParse(joiner.Room)
	var r *room.Room
	if r = room.Get(joiner.RoomID); r == nil {
		//Room Not Found
		fmt.Println("Join Room Error: Room Not Found")
		return errors.New("room not found")
	}

	var user *auth.User
	if user = auth.GetUserByName(joiner.Username); user == nil {
		fmt.Println("Join Room Error: User Not Found")
		return errors.New("user not found")
	}

	if err := r.Join(user); err != nil {
		fmt.Println("Join Room Error: " + err.Error())
		return err
	}

	return nil
}

// ----- Start Game: Initiates a new Game session based on a ready Lobby.
// Only the owner of a Room can start a game, and only if all Players in the Lobby are Ready.
// A new Game instance with its own GUID is generated.
// The existing websocket connections are re-used, with Lobby hooks closed and
// Game channels/goroutines initiated to use the websocket for game state transmission
// func StartGame(c *fiber.Ctx) error {
// 	//TODO: Convert these repeate anon struct declarations into a DTO
// 	starter := struct {
// 		RoomID string `json:"roomID"`
// 		Username string `json:"username"`
// 	}{}

// 	if err := c.BodyParser(&starter); err != nil {
// 		fmt.Println("Game Start Error: " + err.Error())
// 		//TODO: Error/logging pattern
// 		return err
// 	}

// 	room := room.Get(starter.RoomID)
// 	if room == nil {
// 		//Room Not Found
// 		fmt.Println("Game Start Error: Room Not Found")
// 		return errors.New("room not found")
// 	}

// 	user := auth.GetUserByName(starter.Username)
// 	if user == nil {
// 		fmt.Println("Game Start Error: User Not Found")
// 		return errors.New("user not found")
// 	}

// 	if user != room.Owner {
// 		fmt.Println("Game Start Error: Only the Room Owner can start the Game")
// 		return errors.New("not authorized to start game")
// 	}

// 	//Passing all checks, start the game
// 	//First need to make a new Game, as Game is no longer member of Room
// 	g := game.New()
// 	//Iterate over all users in the room lobby and add them as players in the game
// 	for conn, client := range room.Clients() {
// 		g.Add(client.GetUser())
// 		// Need to pass Game ID back to Clients and allow clients to connect to the game
// 		// TODO:
// 	}

// 	// Gives path for authenticated route to Game connection (auth and session ID)
// 	room.Close()
// }
