package game

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/game/grid"
	"github.com/vagrant-technology/squad-leader/messaging"
)

var games = make(map[uuid.UUID]*Game)

type Game struct {
	ID      uuid.UUID     `json:"id"`
	Players map[*auth.User]bool  `json:"players"` //Player Connection Map
	Grid    *grid.HexGrid `json:"grid"`

	hub *messaging.ClientHub `json:"-"`
}

// func (game *Game) ReportState() any {
// 	return game.Grid
// }

func (g *Game) Join(u *auth.User) error {
	if _, exists := g.Players[u]; !exists {
		return errors.New("player not part of game")
	}
}

// func (g *Game) Add(u *auth.User) {
// 	g.Players[u] = false
// }

// func (game *Game) Start() error {

// }

// func (game *Game) NewClient(c *websocket.Conn, user *auth.User) *messaging.Client {

// }

// ----- ----- Static Methods ----- -----

func New(users []*auth.User) string {
	log.Println("Creating a new game instance...")

	game := new(Game)
	game.ID = uuid.New()
	game.Players = make(map[*auth.User]bool)
	game.Grid = grid.NewHexGrid(33, 10)

	for _, u := range users {
		game.Players[u] = false
	}

	//Client-message input processor
	processor := new(GameMessageProcessor)
	game.hub = messaging.NewClientHub(processor.ProcessInput)
	go game.hub.Start()

	games[game.ID] = game

	log.Println("Game " + game.ID.String() + "has been initialized!")
	return game.ID.String()
}

func Get(game string) *Game {
	gameID := uuid.MustParse(game)
	return games[gameID]
}