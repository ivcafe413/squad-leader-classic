package handlers

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/game"
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
		"owner":   user.Username,
	})
}

// ----- Start Game: Initiates a new Game session based on a ready Lobby.
// Only the owner of a Room can start a game, and only if all Players in the Lobby are Ready.
// A new Game instance with its own GUID is generated.
// The existing websocket connections are re-used, with Lobby hooks closed and
// Game channels/goroutines initiated to use the websocket for game state transmission
func CreateGame(c *fiber.Ctx) error {
	//TODO: Convert these repeat anon struct declarations into a DTO
	starter := struct {
		RoomID   string `json:"roomID"`
		Username string `json:"username"`
	}{}

	if err := c.BodyParser(&starter); err != nil {
		//fmt.Println("Game Start Error: " + err.Error())
		log.Println("CreateGame Error: ", err.Error())
		return err
	}

	room := room.Get(starter.RoomID)
	if room == nil {
		//Room Not Found
		// fmt.Println("Game Start Error: Room Not Found")
		err := errors.New("room not found")
		log.Println("CreateGame Error: ", err.Error())
		return err
	}

	user := auth.GetUserByName(starter.Username)
	if user == nil {
		err := errors.New("user not found")
		//fmt.Println("Game Start Error: User Not Found")
		log.Println("CreateGame Error: ", err.Error())
		return err
	}

	if user != room.Owner {
		err := errors.New("not authorized to start game")
		//fmt.Println("Game Start Error: Only the Room Owner can start the Game")
		log.Println("CreateGame Error: ", err.Error())
		return err
	}

	//Passing all checks, start the game
	//Unpack the Users dictionary into a slice of Users
	players := make([]*auth.User, 0, len(room.Users))
	for p := range room.Users {
		players = append(players, p)
	}

	gameID := game.New(players)

	room.Close()

	log.Println("Game " + gameID + " successfully created")
	return c.JSON(&fiber.Map{
		"success": true,
		"gameID":  gameID,
	})
}
