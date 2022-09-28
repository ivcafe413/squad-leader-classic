package game

import (
	//"github.com/vagrant-technology/squad-leader/auth"
	"github.com/vagrant-technology/squad-leader/game/grid"
	"github.com/vagrant-technology/squad-leader/session"
)

type Game struct {
	//Players []*auth.User
	hub *session.ClientHub[*Game]
	grid *grid.HexGrid
}

func (game *Game) ReportState() any {
	return game.grid
}

// -----

func NewGame() *Game {
	game := new(Game)

	game.grid = grid.NewHexGrid(33, 10)
	game.hub = session.NewClientHub(game)
	
	return game
}