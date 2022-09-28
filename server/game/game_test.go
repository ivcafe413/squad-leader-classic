package game

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_GameStateMarshaling(t *testing.T) {
	testGame := NewGame()

	fmt.Println(testGame.grid)
	testMsh, err := json.Marshal(testGame.ReportState())
	if err != nil {
		t.Fatalf("Failure to Marshal Game State: " + err.Error())
	}

	//fmt.Println(testMsh)
	fmt.Println("Marshaled game state: " + string(testMsh))
}