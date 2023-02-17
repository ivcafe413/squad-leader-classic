package game

import (
	"github.com/google/uuid"
	"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/game/grid"
	"github.com/vagrant-technology/squad-leader/messaging"
)

var games = make(map[uuid.UUID]*Game)

type Game struct {
	ID      uuid.UUID           `json:"id"`
	Players map[*auth.User]bool `json:"players"` //Player Connection Map
	Grid    *grid.HexGrid       `json:"grid"`

	hub *messaging.ClientHub `json:"-"`
}

func (game *Game) ReportState() any {
	return game.Grid
}

func (g *Game) Add(u *auth.User) {
	g.Players[u] = false
}

// func (game *Game) Start() error {

// }

// func (game *Game) NewClient(c *websocket.Conn, user *auth.User) *messaging.Client {

// }

// -----

func New() *Game {
	game := new(Game)

	game.ID = uuid.New()
	game.Players = make(map[*auth.User]bool)
	game.Grid = grid.NewHexGrid(33, 10)

	game.hub = messaging.NewClientHub()

	games[game.ID] = game

	return game
}
