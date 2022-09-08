package game

import (
	"github.com/google/uuid"
)

type GameState struct {
	ID uuid.UUID
}

func New() GameState {
	gameID := uuid.New()
	newGame := GameState{
		ID: gameID,
	}

	return newGame
}
