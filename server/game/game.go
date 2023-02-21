package game

import (
	"log"

	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/game/grid"
	"github.com/vagrant-technology/squad-leader/messaging"
)

var games = make(map[uuid.UUID]*Game)

type Game struct {
	ID      uuid.UUID     `json:"id"`
	Players []*auth.User  `json:"players"` //Player Connection Map
	Grid    *grid.HexGrid `json:"grid"`

	hub *messaging.ClientHub `json:"-"`
}

func (game *Game) ReportState() any {
	return game.Grid
}

// func (g *Game) Add(u *auth.User) {
// 	g.Players[u] = false
// }

// func (game *Game) Start() error {

// }

// func (game *Game) NewClient(c *websocket.Conn, user *auth.User) *messaging.Client {

// }

// -----

func New(users []*auth.User) string {
	log.Println("Creating a new game instance...")

	game := new(Game)
	game.ID = uuid.New()
	game.Players = users
	game.Grid = grid.NewHexGrid(33, 10)

	//Client-message input processor
	processor := new(GameMessageProcessor)
	game.hub = messaging.NewClientHub(processor.ProcessInput)
	go game.hub.Start()

	games[game.ID] = game

	log.Println("Game " + game.ID.String() + "has been initialized!")
	return game.ID.String()
}
